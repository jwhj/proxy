FROM scratch
ADD proxy /
EXPOSE 8080
EXPOSE 8081
CMD ["/proxy","-m"]