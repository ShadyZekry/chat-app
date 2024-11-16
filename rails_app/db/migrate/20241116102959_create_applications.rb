class CreateApplications < ActiveRecord::Migration[8.0]
  def change
    create_table :chats, id:false do |t|
      t.string :token, null: false, primary_key: true
      t.string :name

      t.timestamps
    end
  end
end
