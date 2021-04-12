FROM nginx:1.19.2

RUN rm -f /etc/nginx/conf.d/*

COPY conf/nginx.conf /etc/nginx/nginx.conf

CMD /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf
