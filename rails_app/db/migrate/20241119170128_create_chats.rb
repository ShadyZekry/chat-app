class CreateChats < ActiveRecord::Migration[8.0]
  def change
    create_table :chats, id:false do |t|
      t.string :number, null: false, primary_key: true
      t.string :name

      t.timestamps
    end
  end
end
