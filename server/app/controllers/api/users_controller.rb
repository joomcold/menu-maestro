class Api::UsersController < Api::ApiController
  def show
    render(json: { user: current_user }, status: :ok)
  end
end
