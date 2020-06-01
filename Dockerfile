FROM golang:1.14.3
# Get app port from .env
ARG APP_PORT
# Set GO111MODULE to on
# Force using Go modules even if the project is in your GOPATH. Requires go.mod to work
ENV GO111MODULE=on
# Set working dir inside the container
WORKDIR /app
# Copy contents of host to the workdir inside the container
COPY ./ /app
# Runs setup command from bin
RUN /app/bin/cmd/setup
# Exposes the $APP_PORT declared on .env
EXPOSE $APP_PORT
# Runs watch command from bin
CMD ["./bin/cmd/watch"]