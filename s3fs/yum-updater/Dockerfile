FROM centos:7

RUN yum install -y createrepo && yum clean all

VOLUME /data

COPY entrypoint.sh /entrypoint.sh
#ENTRYPOINT ["/entrypoint.sh"]

#CMD [ "createrepo" ]