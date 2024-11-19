class UpdateMessagesCountsJob < ApplicationJob
  queue_as :default

  def perform
    recent_message_counts = Message
      .where('created_at >= ?', 55.minutes.ago)
      .group(:application_token, :chat_number)
      .count

    updates = recent_message_counts.map do |chat_number, messages_count|
      "WHEN number = '#{chat_number}' THEN #{messages_count}"
    end.join(" ")

    Chat.connection.execute(<<~SQL)
    UPDATE chats
    SET messages_count = CASE
      #{updates}
    END
    SQL
  end
end

