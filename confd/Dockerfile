#FROM smebberson/alpine-confd
FROM rancher/confd-base

ADD ./conf.d /etc/confd/conf.d
ADD ./templates /etc/confd/templates

VOLUME /etc/logstash
VOLUME /etc/filebeat
#VOLUME /usr/share/logstash

ENTRYPOINT ["confd"]
CMD ["--backend", "rancher", "--prefix", "/2015-07-25"]
