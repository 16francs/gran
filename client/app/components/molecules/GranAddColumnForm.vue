<template>
  <div class="ma-1">
    <gran-card :width="width" color="grey lighten-3">
      <gran-card-text v-if="isOpen">
        <form>
          <gran-text-field
            v-model="newBoardListForm.name.value"
            :label="newBoardListForm.name.label"
            :rules="newBoardListFormValidate.name"
            autofocus
            @keydown="onKeydown"
          />
          <gran-color-picker
            v-model="newBoardListForm.color.value"
            :label="newBoardListForm.color.label"
          />
          <gran-button
            color="light-blue darken-1"
            :dark="!submitDisabled"
            :disabled="submitDisabled"
            @click="doSubmit"
          >
            Add
          </gran-button>
          <gran-button color="grey darken-1" icon @click="close">
            <gran-icon name="close" />
          </gran-button>
        </form>
      </gran-card-text>
      <v-card-actions v-if="!isOpen">
        <gran-button color="grey darken-1" text block @click="open">
          <gran-icon name="plus" />
          Add Column
        </gran-button>
      </v-card-actions>
    </gran-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import GranCard from '~/components/atoms/GranCard.vue'
import GranCardText from '~/components/atoms/GranCardText.vue'
import GranTextField from '~/components/atoms/GranTextField.vue'
import GranColorPicker from '~/components/atoms/GranColorPicker.vue'
import GranIcon from '~/components/atoms/GranIcon.vue'
import GranButton from '~/components/atoms/GranButton.vue'

import { BoardListForm, BoardListFormValidate } from '~/types/form'

export default Vue.extend({
  components: {
    GranCard,
    GranCardText,
    GranTextField,
    GranColorPicker,
    GranIcon,
    GranButton,
  },
  data: () => {
    return {
      width: 310,
      isOpen: false,
      newBoardListForm: BoardListForm,
      newBoardListFormValidate: BoardListFormValidate,
    }
  },
  computed: {
    submitDisabled(): Boolean {
      return this.newBoardListForm.name.value.length === 0
    },
  },
  methods: {
    open(): void {
      this.isOpen = true
    },
    close(): void {
      this.isOpen = false
    },
    doSubmit(): void {
      if (!this.submitDisabled) {
        this.isOpen = false
        this.$emit('addColumn', this.newBoardListForm)
        this.newBoardListForm.name.value = ''
        this.newBoardListForm.color.value = ''
      }
    },
    onKeydown(keyEvent: KeyboardEvent): void {
      if (keyEvent.keyCode === 13) this.doSubmit() // KeyCode: 13 => enter
      if (keyEvent.keyCode === 27) this.close() // KeyCode: 27 => esc
    },
  },
})
</script>

<style scoped>
.v-btn {
  text-transform: none;
}
</style>
