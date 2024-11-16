Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :chat_applications, only: [:create, :show, :update], path: 'applications', param: :token
    end
  end
end
