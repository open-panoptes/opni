<script>
import { BadgeState } from '@components/BadgeState';

export default {
  components: { BadgeState },
  props:      {
    value: {
      type:     Object,
      required: true
    },
  },

  computed: {
    color() {
      return `bg-${ this.value.state || 'error' }`;
    }
  }
};
</script>

<template>
  <div class="status-badge">
    <BadgeState v-tooltip="value.longMessage" :color="color" :label="value.shortMessage || value.message || 'Unknown'" />
  </div>
</template>

<style lang="scss" scoped>
.status-badge {
  user-select: none;
}
</style>
