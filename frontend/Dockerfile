FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

ADD nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /var/www/html/
ADD css css
ADD images images
ADD js js
ADD static static
ADD index.html index.html

EXPOSE 80