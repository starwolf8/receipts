
FROM node:16.9.1-alpine

# RUN apt-get update && apt-get install android-tools-adb -y
WORKDIR /ui
COPY receipt-ui/package.json receipt-ui/app.json receipt-ui/yarn.lock ./

ADD ./receipt-ui/ /ui

# RUN yarn
RUN yarn global add expo-cli
CMD ["yarn", "web"]

#Reference: 
#   - https://medium.com/@pavan168/pavanwelihinda-running-react-native-in-docker-a0fe0b0c776e