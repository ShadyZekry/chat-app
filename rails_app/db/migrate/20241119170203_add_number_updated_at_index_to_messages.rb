class AddNumberUpdatedAtIndexToMessages < ActiveRecord::Migration[8.0]
  def change
    add_index :messages, [:number, :created_at]
  end
end
