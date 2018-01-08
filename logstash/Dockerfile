FROM docker.elastic.co/logstash/logstash:6.1.1
RUN rm -f /usr/share/logstash/pipeline/logstash.conf
ADD logstash.conf /usr/share/logstash/pipeline/
ADD logstash.yml /usr/share/logstash/config/