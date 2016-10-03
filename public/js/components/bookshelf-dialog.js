Vue.component('bookshelf-dialog', {
    template: '#bookshelf-modal',
    data: function () {
        return {
            title: '',
        };
    },
    methods: {
        save: function () {
            if (this.title) {
                var self = this
                var formData = new FormData();
                formData.append("title", this.title);

                var xhr = new XMLHttpRequest();
                xhr.open("POST", "/bookshelfs");
                xhr.onload = function () {
                    self.$emit('new')
                    self.$emit('close');
                }
                xhr.send(formData);
            }
        }
    }
});