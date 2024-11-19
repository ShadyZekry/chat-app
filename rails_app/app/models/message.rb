class Message < ApplicationRecord
  self.primary_key = "number"

  belongs_to :chat, foreign_key: [:application_token, :chat_number], primary_key: "number"
  validates :name, presence: true
end
