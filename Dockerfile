FROM scratch

WORKDIR /app

ENV NODE_ENV production

COPY staticBuilds/ginFast_linux /app
COPY public /app/public

EXPOSE 9000

CMD ["./ginFast_linux"]