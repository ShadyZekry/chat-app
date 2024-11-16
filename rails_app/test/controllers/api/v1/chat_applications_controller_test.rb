require "test_helper"

class Api::V1::ChatApplicationsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get api_v1_chat_applications_index_url
    assert_response :success
  end

  test "should get show" do
    get api_v1_chat_applications_show_url
    assert_response :success
  end
end
