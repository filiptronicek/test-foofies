from gitpod/workspace-base

COPY ./prebuild.sh .

RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN echo hello
RUN sleep 30

COPY ./not-found.sh .