class AddNumberUpdatedAtIndexToChats < ActiveRecord::Migration[8.0]
  def change
    add_index :chats, [:number, :created_at]
  end
end
