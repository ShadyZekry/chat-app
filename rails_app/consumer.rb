require "bunny"
require "active_record"
require "./config/environment" # Load Rails environment

$stdout.sync = true

TIMEOUT_THRESHOLD = 20 # seconds
BULK_INSERT_THRESHOLD = 5 # queue size

connection = Bunny.new( host: "rabbitmq")
connection.start

chatsChannel = connection.create_channel
chatsQueue = chatsChannel.queue("chats")

messagesChannel = connection.create_channel
messagesQueue = messagesChannel.queue("messages")

# Buffer for messages
chats_buffer = []
messages_buffer = []
last_chats_insert_time = Time.now
last_messages_insert_time = Time.now

def flush_chats_buffer(buffer)
  return if buffer.empty?

  ActiveRecord::Base.transaction do
    buffer.each do |chat|
      Chat.create!(
        number: chat[:number],
        name: chat[:name],
        application_token: chat[:application_token],
      )
    end
  end

  puts "Flushed #{buffer.size} chats to the database."
  buffer.clear
end

def flush_messages_buffer(buffer)
  return if buffer.empty?

  ActiveRecord::Base.transaction do
    buffer.each do |message|
      Message.create!(
        number: message[:number],
        name: message[:name],
        application_token: message[:application_token],
        chat_number: message[:chat_number],
      )
    end
  end

  puts "Flushed #{buffer.size} messages to the database."
  buffer.clear
end

# Periodic flush thread (handles time-based flush)
Thread.new do
  loop do
    sleep 10 # Check every 10 seconds
    if Time.now - last_chats_insert_time > TIMEOUT_THRESHOLD 
      last_chats_insert_time = Time.now
      if chats_buffer.empty?
        puts "Chats buffer is empty, skipping flush."
        next
      end

      puts "Chats time threshold reached. Flushing buffer..."
      flush_chats_buffer(chats_buffer)
    end
  end
end

# Periodic flush thread (handles time-based flush)
Thread.new do
  loop do
    sleep 10 # Check every 10 seconds
    if Time.now - last_messages_insert_time > TIMEOUT_THRESHOLD 
      last_messages_insert_time = Time.now
      if messages_buffer.empty?
        puts "Messages buffer is empty, skipping flush."
        next
      end

      puts "Messages time threshold reached. Flushing buffer..."
      flush_messages_buffer(messages_buffer)
    end
  end
end

Thread.new do
  chatsQueue.subscribe(block: true) do |delivery_info, _properties, body|
    puts "Received chat: #{body}"
    chat = JSON.parse(body)
    chats_buffer << {
      number: chat['number'],
      name: chat['name'],
      application_token: chat['application_token'],
    }

    # Check if buffer size threshold is met
    if chats_buffer.size >= BULK_INSERT_THRESHOLD
      puts "Chats threshold reached. Flushing buffer..."
      flush_chats_buffer(chats_buffer)
    end
  end
end


messagesQueue.subscribe(block: true) do |delivery_info, _properties, body|
  puts "Received message: #{body}"
  chat = JSON.parse(body)
  messages_buffer << {
    number: chat['number'],
    name: chat['name'],
    application_token: chat['application_token'],
    chat_number: chat['chat_number'],
  }

  # Check if buffer size threshold is met
  if messages_buffer.size >= BULK_INSERT_THRESHOLD
    puts "Messages threshold reached. Flushing buffer..."
    flush_messages_buffer(messages_buffer)
  end
end

# Cleanup
at_exit do
  flush_chats_buffer(chats_buffer) unless chats_buffer.empty?
  flush_messages_buffer(messages_buffer) unless messages_buffer.empty?
  connection.close
end

