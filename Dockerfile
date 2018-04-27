FROM ubuntu:16.04

ENV GOPATH /go

# Update the operating system and install base tools:
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927 && \
	# Install specific libssl for mongo-org-* components:
	#wget http://security.debian.org/debian-security/pool/updates/main/o/openssl/libssl1.0.0_1.0.1t-1+deb8u6_amd64.deb && \
	#dpkg -i libssl1.0.0_1.0.1t-1+deb8u6_amd64.deb && \
	# Add mongodb tools:
	echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list && \
	# Install desired components:
	apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y zip mongodb-org-tools mongodb-org-shell git wget && \
	# Create the Go workspace:
	mkdir /go && \
    mkdir /go/src && \
    mkdir /go/bin && \
    mkdir /go/pkg && \
    cd /go && \
	wget --no-check-certificate -O go.tar.gz https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go.tar.gz && \
	rm go.tar.gz

# Install libraries for Re4EEE and Ocean:
RUN export PATH=$PATH:/usr/local/go/bin && \
	# Ocean:
	go get github.com/SommerEngineering/Ocean && \
	# Generator for UUIDs:
	go get github.com/twinj/uuid && \
	# Database driver:
	go get gopkg.in/mgo.v2 

# Insert all files from the repo (but from the current directory, not from Git):
ADD . /go/src/github.com/SommerEngineering/Re4EEE/

# Compile and Setup
RUN	export PATH=$PATH:/usr/local/go/bin && \
	cd /go/src/github.com/SommerEngineering/Re4EEE && \
	# Compile the Re4EEE:
	go install && \
	# Copy the final binary and the runtime scripts to the home folder:
	cp /go/bin/Re4EEE /home && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/run.sh /home/run.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/setConfiguration.sh /home/setConfiguration.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/configureCustomerDB.sh /home/configureCustomerDB.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/uploadStaticData.sh /home/uploadStaticData.sh && \
	# Zip static data and move them to the home folder:
	cd staticFiles && \
	zip -r /home/staticFiles.zip . && \
	cd ../templates && \
	zip -r /home/templates.zip . && \
	cd ../web && \
	zip -r /home/web.zip . && \
	# Uninstall tools:
	apt-get autoremove -y zip && \
	# Delete the entire Go workspace:
	rm -r -f /go && \
	# Create the configuration file:
	touch /home/configuration.json && \
	touch /home/project.name && \
	# Make the scripts executable:
	chmod 0777 /home/run.sh && \
	chmod 0777 /home/setConfiguration.sh && \
	chmod 0777 /home/configureCustomerDB.sh && \
	chmod 0777 /home/uploadStaticData.sh && \
	chmod 0777 /home/Re4EEE && \
	chmod 0666 /home/configuration.json && \
	chmod 0666 /home/project.name

# Run anything below as nobody:
USER nobody

# Re4EEE provides HTTP by port 40000 and the admin interface on port 50000:
EXPOSE 40000 50000

# Define the working directory:
WORKDIR /home

# The default command to run, if a container starts:
CMD ["./run.sh"]
