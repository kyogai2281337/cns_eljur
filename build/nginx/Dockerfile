# Используем nginx в качестве базового образа для этапа production
FROM nginx

# Копируем собранное приложение из предыдущего этапа в папку /usr/share/nginx/html внутри контейнера
COPY frontend/cns_eljur/dist /usr/share/nginx/html

COPY build/nginx/nginx.conf /etc/nginx/nginx.conf