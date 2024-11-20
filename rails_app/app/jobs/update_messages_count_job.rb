class UpdateMessagesCountsJob < ApplicationJob
  queue_as :default

  def perform
    recent_message_counts = Message
      .where('created_at >= ?', 55.minutes.ago)
      .group(:application_token, :chat_number)
      .count

    scope = recent_message_counts.map do |chat_id, messages_count|
      "concat('#{chat_id[0]}', '#{chat_id[1]}')"
    end.join(", ")

    updates = recent_message_counts.map do |chat_id, messages_count|
      "WHEN concat('#{chat_id[0]}', '#{chat_id[1]}') THEN messages_count + #{messages_count}"
    end.join(" ")

    Chat.connection.execute(<<~SQL)
    UPDATE chats
    SET messages_count = CASE concat(application_token, number)
      #{updates}
    END
    where concat(application_token, number) IN (#{scope})
    SQL
  end
end

