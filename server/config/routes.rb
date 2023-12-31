Rails.application.routes.draw do
  devise_for :users, path: '',
                     path_names: {
                       sign_in: 'login',
                       sign_out: 'logout',
                       registration: 'signup'
                     },
                     controllers: {
                       sessions: 'api/users/sessions',
                       registrations: 'api/users/registrations'
                     }

  get 'profile', to: 'api/users#show'
end
