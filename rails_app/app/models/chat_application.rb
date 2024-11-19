class ChatApplication < ApplicationRecord
  self.table_name = "applications"
  self.primary_key = "token"

  has_many :chats, foreign_key: "application_token", primary_key: "token"
  validates :name, presence: true
end
