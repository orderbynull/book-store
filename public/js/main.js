var SHELFS = "/bookshelfs/";
var BOOKS = "/books/";
var AUTHORS = "/authors/";

var vm = new Vue({
    el: '#app',
    data: {
        showBookshelfModal: false,
        showBookModal: false,
        activeBookshelf: 0,
        bookshelfsLoading: false,
        booksLoading: false,

        bookshelfs: [],
        books: [],
        authors: []
    },

    methods: {
        changeBookshelf: function (id) {
            this.activeBookshelf = id;
            this.loadBooks(id);
        },

        loadBookshelfs: function () {
            this.activeBookshelf = 0;
            this.bookshelfsLoading = true;

            var xhr = new XMLHttpRequest()
            var self = this
            xhr.open('GET', SHELFS)
            xhr.onload = function () {
                json = JSON.parse(xhr.responseText)
                if(!json)
                    return;

                self.bookshelfs = json
                self.bookshelfsLoading = false;

                if(self.bookshelfs && self.bookshelfs.length > 0){
                    self.activeBookshelf = self.bookshelfs[0].id;
                    self.loadBooks()
                }
            }
            xhr.send()
        },

        loadBooks: function () {
            if(this.activeBookshelf){
                this.booksLoading = true;

                var xhr = new XMLHttpRequest()
                var self = this
                xhr.open('GET', BOOKS + this.activeBookshelf)
                xhr.onload = function () {
                    self.books = JSON.parse(xhr.responseText)
                    self.booksLoading = false;
                }
                xhr.send()
            }
        },

        deleteBook: function (id) {
            if (confirm("Are you sure?")) {
                var xhr = new XMLHttpRequest()
                var self = this
                xhr.open('DELETE', BOOKS + id)
                xhr.onload = function () {
                    self.loadBooks();
                }
                xhr.send()
            }
        },
        
        loadAuthors: function () {
            var xhr = new XMLHttpRequest()
            var self = this
            xhr.open('GET', AUTHORS)
            xhr.onload = function () {
                self.authors = JSON.parse(xhr.responseText)
            }
            xhr.send()
        }
    }
})

vm.loadAuthors();
vm.loadBookshelfs();