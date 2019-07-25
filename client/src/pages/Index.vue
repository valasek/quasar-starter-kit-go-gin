<template>
  <q-page>
    <div class="row justify-center">
      <img alt="Quasar logo" src="~assets/quasar-logo-full.svg">
    </div>
    <h4 class="row justify-center">
      Ready to Use Scenarios
    </h4>
    <div class="row justify-center q-gutter-lg">
      <div class="column">
        <div class="text-subtitle2 q-pa-sm">Download file from the server</div>
        <q-btn color="primary" @click="download">
          Download&nbsp;&nbsp;<q-icon name="cloud_download"></q-icon>
        </q-btn>
      </div>
      <div class="column">
        <div class="text-subtitle2 q-pa-sm">Upload files to the server and allow only zip file</div>
          <q-uploader
            label="Upload"
            accept=".zip"
            :url=uploaderUrl
            auto-upload
            style="max-width: 250px"
          />
      </div>
    </div>
  </q-page>
</template>

<script>
import api from '../api/axiosSettings'

export default {
  name: 'PageIndex',

  data () {
    return {
      uploaderUrl: 'http://localhost:3000/api/upload/data',
      uploadFile: {}
    }
  },

  methods: {
    upload (file) {
      const comp = this
      const fr = new FileReader()
      fr.readAsDataURL(file)
      fr.addEventListener('load', () => {
        // this.uploadedImage = fr.result
        this.uploadFile = file
        const formData = new FormData()
        formData.append('uploadFile', this.uploadFile)
        api.apiClient.post('/api/upload/data', formData)
          .then(response => {
            this.$q.notify({
              message: 'Restore successfully completed',
              color: 'teal'
            })
            // implement post upload steps if required
          })
          .catch(function (e) {
            comp.$q.notify({
              message: 'upload successfully completed',
              color: 'negative',
              icon: 'report_problem'
            })
            console.log(e, e.response) /* eslint-disable-line no-console */
          })
      })
    },
    download () {
      const comp = this
      api.apiClient.get('/api/download/data', { responseType: 'blob' })
        .then(response => {
          const url = window.URL.createObjectURL(new Blob([response.data]))
          const link = document.createElement('a')
          link.href = url
          link.setAttribute('download', 'README.md')
          document.body.appendChild(link)
          link.click()
        })
        .catch(function (e) {
          comp.$q.notify({
            message: 'Couldn\'t download data\n' + e.toString(),
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e) /* eslint-disable-line no-console */
        })
    }
  }
}
</script>

<style scoped>
</style>
