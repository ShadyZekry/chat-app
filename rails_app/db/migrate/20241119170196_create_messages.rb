class CreateMessages < ActiveRecord::Migration[8.0]
  def change
    create_table :messages, id:false do |t|
      t.string :number, null: false, primary_key: true
      t.string :name
      t.integer :chat_number, null: false
      t.string :application_token, null: false

      t.timestamps
    end
    add_foreign_key :messages, :chats, column: [:application_token, :chat_number], primary_key: [:application_token, :number], index: true
    add_index :messages, [:chat_number, :created_at]
  end
end
