<template>
  <div class="notifications-container">
    <transition-group name="notification">
      <div v-for="(notification, index) in notifications" :key="notification.id" class="notification">
        <div class="notification-content" @click="hideNotification(index)">
          {{ notification.message }}
        </div>
        <div class="progress-bar" :style="{ width: notification.progress + '%' }"></div>
      </div>
    </transition-group>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

interface Notification {
  id: number;
  message: string;
  progress: number;
}

export default defineComponent({
  name: 'NotificationsUi',
  setup() {
    const notifications = ref<Notification[]>([]);

    const addNotification = (message: string) => {
      const newNotification: Notification = {
        id: Date.now(),
        message,
        progress: 100,
      };
      notifications.value.push(newNotification);
      setTimeout(() => {
        hideNotification(notifications.value.indexOf(newNotification));
      }, 5000);
    };

    const hideNotification = (index: number) => {
      notifications.value.splice(index, 1);
    };

    return {
      notifications,
      addNotification,
      hideNotification,
    };
  },
});
</script>

<style scoped>
.notifications-container {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 1000;
}

.notification {
  margin-bottom: 10px;
  background-color: #ffffff;
  border: 1px solid #cccccc;
  border-radius: 4px;
  padding: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.notification-content {
  cursor: pointer;
}

.progress-bar {
  height: 4px;
  background-color: #007bff;
  transition: width 0.5s ease-in-out;
  border-radius: 2px;
}
</style>
