# Base image
FROM ruby:3.2

# Install dependencies
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs cron

# Create app directory
WORKDIR /app

# Install gems
COPY Gemfile* ./
RUN bundle install

# Copy the rest of the application code
COPY . .

# Expose Rails port
EXPOSE 3000

RUN groupadd -g 1000 appuser && useradd -m -u 1000 -g appuser appuser
USER appuser

RUN whenever --update-crontab

# Start the Rails server
# CMD ["rails", "server", "-b", "0.0.0.0"]
CMD ["ruby", "-S", "rails", "server", "-b", "0.0.0.0"]
