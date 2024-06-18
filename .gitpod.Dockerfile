FROM gitpod/workspace-base:latest

ENV TRIGGER_REBUILD=52

RUN bash -c "for i in $(seq -s " " 1 120); do echo    "$i"; sleep 1; done"
