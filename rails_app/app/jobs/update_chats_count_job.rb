class UpdateChatsCountsJob < ApplicationJob
  queue_as :default

  def perform
    recent_chats_counts = Chat
      .where('created_at >= ?', 55.minutes.ago)
      .group(:application_token)
      .count

    updates = recent_chats_counts.map do |application_token, chats_count|
      "WHEN token = '#{application_token}' THEN #{chats_count}"
    end.join(" ")

    ChatApplication.connection.execute(<<~SQL)
    UPDATE applications
    SET chats_count = CASE
      #{updates}
    END
    SQL
  end
end

