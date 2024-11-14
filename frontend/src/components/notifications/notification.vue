<template>
  <div class="notification-container" v-if="notifications.length">
    <transition-group name="notification-fade">
      <div
        v-for="notification in notifications"
        :key="notification.id"
        :class="['notification', `notification-${notification.type}`]"
      >
        <h4 class="h4">{{ notification.title }}</h4>
        <p class="text">{{ notification.message }}</p>
        <button
          class="notification-close btn-no-bg"
          @click="removeNotification(notification.id)"
        >
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M6 6L18 18"
              stroke="#A6A6A6"
              stroke-width="2"
              stroke-linecap="round"
            />
            <path
              d="M18 6L6 18"
              stroke="#A6A6A6"
              stroke-width="2"
              stroke-linecap="round"
            />
          </svg>
        </button>
      </div>
    </transition-group>
  </div>
</template>

<script>
export default {
  name: "Notification",
  data() {
    return {
      notifications: [],
    };
  },
  methods: {
    addNotification(notification) {
      const id = Date.now();
      this.notifications.push({
        id,
        title: notification.title || "Уведомление",
        message: notification.message,
        type: notification.type || "default",
        timeout: notification.timeout || 5000,
      });

      if (notification.timeout !== 0) {
        setTimeout(() => {
          this.removeNotification(id);
        }, notification.timeout);
      }
    },
    removeNotification(id) {
      this.notifications = this.notifications.filter((item) => item.id !== id);
    },
    clearAll() {
      this.notifications = [];
    },
  },
  created() {
    this.$root.$on("add-notification", this.addNotification);
    this.$root.$on("clear-notifications", this.clearAll);
  },
  beforeDestroy() {
    this.$root.$off("add-notification", this.addNotification);
    this.$root.$off("clear-notifications", this.clearAll);
  },
};
</script>

<style scoped>
.notification-container {
  position: fixed;
  top: 1em;
  right: 1em;
  z-index: 9999;
  max-width: 17.5em;
}

.notification {
  position: relative;
  padding: 0.7em 2em 0.7em 0.7em;
  margin-bottom: 10px;
  border-radius: 4px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.notification-close {
  position: absolute;
  top: .5em;
  right: .5em;
  cursor: pointer;
}

.notification-default {
  border-left: 4px solid #a6a6a6;
}

.notification-info {
  border-left: 4px solid #2196f3;
}

.notification-warning {
  border-left: 4px solid #ff9800;
}

.notification-succes {
  border-left: 4px solid #4caf50;
}

.notification-fade-enter-active,
.notification-fade-leave-active {
  transition: all 0.3s ease;
}

.notification-fade-enter,
.notification-fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
