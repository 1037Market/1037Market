<template>
  <div>
    <vue-advanced-chat
        height="100vh"
        :current-user-id="currentUserId"
        :rooms="rooms"
        :rooms-loaded="true"
        :messages="messages"
        :messages-loaded="messagesLoaded"
        :room-id="currentRoomId"
        @send-message="sendMessage($event.detail[0])"
        @fetch-messages="fetchMessages($event.detail[0])"
        @fetch-more-rooms="fetchMoreRooms"
        @roomChanged="roomChanged"
    >


        <!--      可以加slot-->

    </vue-advanced-chat>
  </div>
</template>

<script>
import { register } from 'vue-advanced-chat'
import {getMessageContent, getNMessages, getRoomIds, getRoomInfo, sendMessage} from "../../network/chat";
import {nextTick} from "vue";
import {uploadImage} from "../../network/image";
register()

async function convertBlobUrlToFile(blobUrl, fileName, mimeType) {
  // 使用 fetch() 来获取 Blob 数据
  const response = await fetch(blobUrl);
  const blob = await response.blob();

  // 创建一个 File 对象
  const file = new File([blob], fileName, { type: mimeType });

  return file;
}

export default {
  name: 'Chat',
  data() {
    return {
      currentUserId: window.localStorage.getItem('studentId'),
      rooms: [],
      messages: [],
      messagesLoaded: true,
      currentRoomId: null,
      msgMap: {}
    }
  },
  methods: {
    fetchMessages(options) {
        let room = options.room;
        if(options.options.reset) {
          this.msgMap[room.roomId] = [];
          this.messages = this.msgMap[room.roomId];
          let alreadyHave = this.messages.length;
          console.log('already heve: ', alreadyHave)
          getNMessages(room.roomId, alreadyHave + 1, 100).then((messageIds) => {
            messageIds.forEach((msgId) => {
              getMessageContent(msgId).then((msgDetail) => {
                let msg = {
                  _id: msgDetail.messageId,
                  indexId: msgDetail.messageId,
                  content: msgDetail.content,
                  senderId: msgDetail.fromId,
                  files: msgDetail.imageURI === '' ? null : [
                    {
                      name: 'My File',
                      type: 'png',
                      audio: false,
                      duration: 14.4,
                      url: 'http://franky.pro:7301/api/image?imageURI=' + msgDetail.imageURI,
                    }
                  ],
                  timestamp: msgDetail.sendTime,
                  date: new Date().toDateString()
                };
                this.messages = [...this.messages, msg]
              })
            })
          }).catch((err) => {
            console.log(err);
          })
        } else {
          console.log('no reset')
          this.messages = this.msgMap[room.roomId];
        }
    },


    sendMessage(message) {
      console.log(message)
      if(message.files) {
        convertBlobUrlToFile(message.files[0].localUrl, message.files[0].name, message.files[0].type).then((file) => {
          let formData = new FormData();
          formData.append('file', file)
          uploadImage(formData).then((response) => {
            let imageURI = response;
            let msg = {
              sessionId: message.roomId,
              content: message.content,
              imageURI: imageURI
            }
            sendMessage(msg).then((response) => {
              msg = {
                _id: response,
                indexId: response,
                content: message.content,
                senderId: this.currentUserId,
                files: [
                  {
                    name: 'My File',
                    type: 'png',
                    audio: false,
                    duration: 14.4,
                    url: 'http://franky.pro:7301/api/image?imageURI=' + msg.imageURI
                  }
                ],
                timestamp: new Date().toString().substring(16, 21),
                date: new Date().toDateString()
              }
              this.messages = [...this.messages, msg]
            }).catch((err) => {
              console.log(err)
            })

          }).catch((err) => {
            console.log(err);
          })
        }).catch((err) => {
          console.log(err);
        })

      } else {
        let msg = {
          sessionId: message.roomId,
          content: message.content,
          imageURI: ''
        }
        sendMessage(msg).then((response) => {
          msg = {
            _id: response,
            indexId: response,
            content: message.content,
            senderId: this.currentUserId,
            files: null,
            timestamp: new Date().toString().substring(16, 21),
            date: new Date().toDateString()
          }
          this.messages = [...this.messages, msg]
        }).catch((err) => {
          console.log(err)
        })
      }


    },

    fetchMoreRooms() {
      console.log('fetch rooms')
    },
    roomChanged(newRoomId) {
      console.log(newRoomId)
    },

  },
  mounted() {
    getRoomIds().then((response) => {
      let rooms = []
      let oldRooms = this.rooms;
      response.forEach((roomId) => {
        getRoomInfo(roomId).then((roomInfo) => {
          let idx = roomInfo[0].userId == window.localStorage.getItem('studentId') ? 1 : 0;
          rooms.push({
            roomId: roomId,
            roomName: roomInfo[idx].nickName,
            avatar: 'http://franky.pro:7301/api/image?imageURI=' + roomInfo[idx].avatar,
            users: [
              {_id: roomInfo[0].userId, username: roomInfo[0].nickName},
              {_id: roomInfo[1].userId, username: roomInfo[1].nickName},
            ]
          })
          this.rooms = [...oldRooms, ...rooms]
          console.log(this.rooms)
        })
      })

    })

  },
  watch: { // 永远按顺序排列
    'messages'(oldVal, newVal) {
      this.messages.sort((x, y) => {
        return x._id - y._id;
      })
    }
  }
}
</script>

<style lang="scss">
body {
  font-family: 'Quicksand', sans-serif;
}
</style>