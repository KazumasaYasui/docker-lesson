FROM nginx:1.19.2

RUN rm -f /etc/nginx/conf.d/*

COPY ./conf/nginx.conf.template /etc/nginx/conf.d/app.conf.template

CMD /bin/sh -c "envsubst < /etc/nginx/conf.d/app.conf.template > /etc/nginx/conf.d/app.conf && nginx -g 'daemon off;'"
