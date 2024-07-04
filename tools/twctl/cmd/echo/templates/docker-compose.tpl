services:
  app:
    build: 
      context: .
      target: dev
    ports:
      - 8888:8888
    env_file:
      - .env
    environment:
      - GO_ENV=production
    privileged: true
    volumes:
      - ./:/app
      - ./db:/db
    restart: always
    logging:
      driver: "json-file"
      options:
        max-size: "10m" # Maximum size of the log file before it gets rotated
        max-file: "3"   # Maximum number of log files to keep
    networks:
      - {{.serviceName}}

networks:
  {{.serviceName}}:
    driver: bridge
