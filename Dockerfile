FROM scratch
ADD main /
ADD public /public
EXPOSE 8080
CMD ["/main"]