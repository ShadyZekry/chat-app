class ChatApplication < ApplicationRecord
  self.table_name = "applications"
  self.primary_key = "token"

  validates :name, presence: true
end
