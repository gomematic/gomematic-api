FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Gomematic API" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 8080 8090
VOLUME ["/var/lib/gomematic"]

ENTRYPOINT ["/usr/bin/gomematic-api"]
CMD ["server"]

ENV GOMEMATIC_API_DB_DSN boltdb:///var/lib/gomematic/database.db
ENV GOMEMATIC_API_UPLOAD_DSN file:///var/lib/gomematic/uploads

RUN apk add --no-cache ca-certificates mailcap bash

COPY dist/binaries/gomematic-api-*-linux-amd64 /usr/bin/
