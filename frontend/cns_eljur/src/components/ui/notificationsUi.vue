<template>
  <div class="notifications-container">
    <transition-group name="notification">
      <div
        v-for="(notification, index) in notifications"
        :key="notification.id"
        :data-notification-id="notification.id"
        class="notification"
      >
        <div class="notification-content" @click="hideNotification(index)">
          {{ notification.message }}
        </div>
        <div class="progress-bar" :style="{ width: notification.progress + '%' }"></div>
      </div>
    </transition-group>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue';

interface Notification {
  id: number;
  message: string;
  progress: number;
  intervalId?: number;
  isClosed?: boolean;
}

export default defineComponent({
  name: 'NotificationsUi',
  setup() {
    const notifications = reactive<Notification[]>([]);

    const addNotification = (message: string) => {
      const newNotification: Notification = {
        id: Date.now(),
        message,
        progress: 0,
        isClosed: false,
      };
      notifications.push(newNotification);

      newNotification.intervalId = setInterval(() => {
        if (newNotification.isClosed) {
          clearInterval(newNotification.intervalId);
          return;
        }

        newNotification.progress += 1;
        const element = document.querySelector(`[data-notification-id="${newNotification.id}"] .progress-bar`) as HTMLElement;
        if (element) {
          element.style.width = newNotification.progress + '%';
        }

        if (newNotification.progress >= 100) {
          clearInterval(newNotification.intervalId);
          hideNotification(notifications.indexOf(newNotification));
        }
      }, 50);
    };

    const hideNotification = (index: number) => {
      const notification = notifications[index];
      if (notification) {
        notification.isClosed = true;
        notifications.splice(index, 1);
      }
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
  transition: width 0.05s ease-in-out;
  border-radius: 2px;
}
</style>
