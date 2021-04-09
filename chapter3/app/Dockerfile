FROM nginx:1.19.2

# ENV APP_SECRET_DATA="hogehoge"

RUN rm -f /etc/nginx/conf.d/*

COPY ./conf/nginx.conf.template /etc/nginx/conf.d/app.conf.template

# EXPOSE 80

CMD /bin/sh -c "envsubst < /etc/nginx/conf.d/app.conf.template > /etc/nginx/conf.d/app.conf && nginx -g 'daemon off;'"
