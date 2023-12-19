FROM alpine:latest

WORKDIR /
COPY "testwebapp" "/testwebapp"
ENV PORT=3001
ENV APP_ID=app1
RUN "ls"
CMD [ "/testwebapp" ]