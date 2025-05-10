class Api::V1::ChatApplicationsController < ApplicationController
  def show
    application = ChatApplication.find_by(token: params[:token])
    if application
      render json: application, status: :ok
    else
      render json: {error: 'Not found'}, status: :not_found
    end
  end

  def create
    application = ChatApplication.new(application_params)
    application.token = SecureRandom.uuid
    if application.save
      add_app_to_redis(application.token)
      render json: {token: application.token, name: application.name}, status: :created
    else
      render json: {error: application.errors.full_messages}, status: :unprocessable_entity
    end
  end

  def update
    application = ChatApplication.find_by(token: params[:token])
    if application
      if application.update(application_params)
        render json: {token: application.token, name: application.name}, status: :ok
      else
        render json: {error: application.errors.full_messages}, status: :unprocessable_entity
      end
    else
      render json: {error: 'Not found'}, status: :not_found
    end
  end

  def application_params
    params.permit(:name)
  end

  def add_app_to_redis(token)
    redis = Redis.new(url: ENV['REDIS_URL'])
    redis.hset(token, 'chats', '0')
  end
end
