require "bunny"
require "active_record"
require "./config/environment" # Load Rails environment

$stdout.sync = true

TIMEOUT_THRESHOLD = 20 # seconds
BULK_INSERT_THRESHOLD = 5 # queue size

connection = Bunny.new( host: "rabbitmq")
connection.start
channel = connection.create_channel
queue = channel.queue("chats")

# begin
#   puts ' [*] Waiting for messages. To exit press CTRL+C'
#   queue.subscribe(manual_ack: false, block: true) do |delivery_info, properties, body|
#     chat_count = channel.queue_declare(queue.name, passive: true).message_count
#     puts " [x] Received #{body} count: #{chat_count}"
#     # data = JSON.parse(body)
#     # ChatMessage.create!(chat_id: data["chat_id"], message_id: data["message_id"], content: data["content"])
#   end
# rescue Interrupt => _
#   connection.close
#
#   exit(0)
# end







# Buffer for messages
chats_buffer = []
last_insert_time = Time.now

# Function to flush buffer to MySQL
def flush_buffer(buffer)
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

  puts "Flushed #{buffer.size} messages to the database."
  buffer.clear
end

# Periodic flush thread (handles time-based flush)
Thread.new do
  loop do
    sleep 10 # Check every 10 seconds
    if Time.now - last_insert_time > TIMEOUT_THRESHOLD 
      last_insert_time = Time.now
      if chats_buffer.empty?
        puts "Buffer is empty, skipping flush."
        next
      end

      puts "Time threshold reached. Flushing buffer..."
      flush_buffer(chats_buffer)
    end
  end
end

# Subscribe to RabbitMQ
queue.subscribe(block: true) do |delivery_info, _properties, body|
  puts "Received message: #{body}"
  chat = JSON.parse(body)
  chats_buffer << {
    number: chat['number'],
    name: chat['name'],
    application_token: chat['application_token'],
  }

  # Check if buffer size threshold is met
  if chats_buffer.size >= BULK_INSERT_THRESHOLD
    puts "Message threshold reached. Flushing buffer..."
    flush_buffer(chats_buffer)
  end
end

# Cleanup
at_exit do
  flush_buffer(chats_buffer) unless chats_buffer.empty?
  connection.close
end

