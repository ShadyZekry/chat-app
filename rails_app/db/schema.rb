# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.0].define(version: 2024_11_19_180204) do
  create_table "applications", primary_key: "token", id: :string, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "name"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "chats_count", default: 0, null: false
  end

  create_table "chats", primary_key: ["application_token", "number"], charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.integer "number", null: false
    t.string "name"
    t.string "application_token", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "messages_count", default: 0, null: false
    t.index ["application_token", "created_at"], name: "index_chats_on_application_token_and_created_at"
  end

  create_table "messages", primary_key: "number", id: :string, charset: "utf8mb4", collation: "utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "name"
    t.integer "chat_number", null: false
    t.string "application_token", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["application_token", "chat_number"], name: "fk_rails_98a3768c6c"
    t.index ["chat_number", "created_at"], name: "index_messages_on_chat_number_and_created_at"
  end

  add_foreign_key "chats", "applications", column: "application_token", primary_key: "token"
  add_foreign_key "messages", "chats", column: ["application_token", "chat_number"], primary_key: ["application_token", "number"]
end
