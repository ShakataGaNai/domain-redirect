# Use an official Python runtime as a parent image
FROM python:3.12-slim

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the requirements file into the container at /usr/src/app
# (Assuming you have a requirements.txt file that includes Flask)
COPY requirements.txt ./

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the current directory contents into the container at /usr/src/app
COPY . .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV REDIRECT_DOMAIN=example.com
ENV LISTEN_ADDR=0.0.0.0:8080

# Run app.py when the container launches
CMD ["python", "./main.py"]