<template>

  <van-nav-bar title="联系卖家" fixed placeholder style="--van-nav-bar-background: rgba(66, 185, 131, 0.9);--van-nav-bar-title-text-color: rgba(255,255,255,1);"
               left-arrow @click-left="router.go(-1)"
  />

  <div>
    <vue-advanced-chat
        :height="needNav ? '95vh' : '89vh'"
        :current-user-id="currentUserId"
        :rooms="rooms"
        :rooms-loaded="true"
        :messages="messages"
        :messages-loaded="messagesLoaded"
        :single-room="needNav"
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
import {
  getAllMsgs,
  getMessageContent,
  getNewestMessageIn,
  getNMessages,
  getRoomIds,
  getRoomInfo,
  sendMessage
} from "../../network/chat";
import {nextTick, onUnmounted} from "vue";
import {uploadImage} from "../../network/image";
import {useRoute} from "vue-router/dist/vue-router";
import {useRouter} from "vue-router";
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
      msgMap: {},
      needNav: false,
      router: useRouter()
    }
  },
  methods: {
    fetchMessages(options) {
        let room = options.room;
        this.currentRoomId = room.roomId;
        if(options.options.reset) {
          this.msgMap[room.roomId] = [];
          this.messages = this.msgMap[room.roomId];
          getAllMsgs(room.roomId).then((response) => {
            response.forEach((msg) => {
              let mes = {
                _id: msg.messageId,
                indexId: msg.messageId,
                content: msg.content,
                senderId: msg.fromId,
                files: msg.imageURI === "" ? null : [
                  {
                    name: 'My File',
                    type: 'png',
                    audio: false,
                    duration: 14.4,
                    url: 'http://franky.pro:7301/api/image?imageURI=' + msg.imageURI
                  }
                ],
                timestamp: msg.sendTime,
                date: new Date().toDateString()
              }
              this.messages = [...this.messages, mes]
            })
          }).catch((err) => {
            console.log(err);
          })
        } else {
          this.messages = this.msgMap[room.roomId];
        }
    },


    sendMessage(message) {
      if(message.files) {
        convertBlobUrlToFile(message.files[0].localUrl, message.files[0].name, message.files[0].type).then((file) => {
          let formData = new FormData();
          formData.append('file', file)
          uploadImage(formData).then((response) => {
            let msg = {
              sessionId: message.roomId,
              content: message.content,
              imageURI: response
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
    const route = useRoute();
    let sessionId = route.params.session;
    if(sessionId === "0") {
      this.needNav = false;
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
          })
        })
      })
    } else {
      this.needNav = true;
      const nav = document.getElementById('nav');
      nav.style.visibility = 'hidden';
      getRoomInfo(sessionId).then((roomInfo) => {
        let idx = roomInfo[0].userId == window.localStorage.getItem('studentId') ? 1 : 0;
        this.rooms = [{
          roomId: Number(sessionId),
          roomName: roomInfo[idx].nickName,
          avatar: 'http://franky.pro:7301/api/image?imageURI=' + roomInfo[idx].avatar,
          users: [
            {_id: roomInfo[0].userId, username: roomInfo[0].nickName},
            {_id: roomInfo[1].userId, username: roomInfo[1].nickName},
          ]
        }]
      })
    }


    let interval = setInterval(() => {
      if(this.currentRoomId !== null) {
        getNewestMessageIn(this.currentRoomId).then((response) => {
          if(response !== this.messages[this.messages.length - 1]._id) {
            let lastId = this.messages[this.messages.length - 1]._id;
            getAllMsgs(this.currentRoomId).then((response) => {
              response.forEach((msg) => {
                let mes = {
                  _id: msg.messageId,
                  indexId: msg.messageId,
                  content: msg.content,
                  senderId: msg.fromId,
                  files: msg.imageURI === "" ? null : [
                    {
                      name: 'My File',
                      type: 'png',
                      audio: false,
                      duration: 14.4,
                      url: 'http://franky.pro:7301/api/image?imageURI=' + msg.imageURI
                    }
                  ],
                  timestamp: msg.sendTime,
                  date: new Date().toDateString()
                }
                if(mes._id > lastId) {
                  this.messages = [...this.messages, mes]
                }
              })
            }).catch((err) => {
              console.log(err);
            })
          }
        }).catch((err) => {
          console.log(err);
        })
      }

    }, 3000)
    onUnmounted(() => {
      clearInterval(interval);
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