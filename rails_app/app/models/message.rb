class Message < ApplicationRecord
  self.primary_key = [:application_token, :chat_number, :number]

  belongs_to :chat, foreign_key: [:application_token, :chat_number], primary_key: [:application_token, :number]
  validates :name, presence: true
end
