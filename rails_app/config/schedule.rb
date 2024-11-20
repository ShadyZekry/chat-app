every 55.minutes do
  runner "UpdateChatsCountJob.perform_now"
  runner "UpdateMessagesCountJob.perform_now"
end
