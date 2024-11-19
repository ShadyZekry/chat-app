class CreateChats < ActiveRecord::Migration[8.0]
  def change
    create_table :chats, primary_key: [:application_token, :number] do |t|
      t.integer :number, null: false
      t.string :name
      t.string :application_token, null: false

      t.timestamps
    end
    add_foreign_key :chats, :applications, column: :application_token, primary_key: :token, index: true
    add_index :chats, [:application_token, :created_at]
  end
end
