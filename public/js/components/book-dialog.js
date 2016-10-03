Vue.component('book-dialog', {
    template: '#book-modal',
    props: ['authors', 'bookshelf'],
    data: function () {
        return {
            title: '',
            author_id: 0
        };
    },
    methods: {
        close: function () {
            this.show = false;
        },
        save: function () {
            if(this.bookshelf && this.title && this.author_id){
                var self = this
                var formData = new FormData();
                formData.append("title", this.title);
                formData.append("author_id", this.author_id);
                formData.append("bookshelf_id", this.bookshelf);

                var xhr = new XMLHttpRequest();
                xhr.open("POST", "/books");
                xhr.onload = function () {
                    self.$emit('new')
                    self.$emit('close');
                }
                xhr.send(formData);
            }
        },
    }
});