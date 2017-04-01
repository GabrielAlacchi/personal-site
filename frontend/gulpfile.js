var gulp = require('gulp');
var plumber = require('gulp-plumber');
var uglify = require('gulp-uglify');
var sass = require('gulp-sass');
var rename = require('gulp-rename');
var autoprefixer = require('gulp-autoprefixer');

var browserSync = require('browser-sync').create();

gulp.task('scripts', function() {
  return gulp.src('js/*.js')
    .pipe(plumber(plumber({
      errorHandler: function (err) {
        console.log(err);
        this.emit('end');
      }
    })))
    .pipe(uglify({
      preserveComments: 'license'
    }))
    .pipe(rename({extname: '.min.js'}))
    .pipe(gulp.dest('js'))
    .pipe(browserSync.stream());
});

gulp.task('styles', function() {
  return gulp.src('scss/*.scss')
    .pipe(plumber(plumber({
      errorHandler: function (err) {
        console.log(err);
        this.emit('end');
      }
    })))
    .pipe(sass({outputStyle: 'compressed'}))
    .pipe(autoprefixer())
    .pipe(gulp.dest('www/css'))
    .pipe(browserSync.stream());
});

gulp.task('update', function() {
  return gulp.src('www/*.html')
    .pipe(browserSync.stream());
});

gulp.task('watch', ['scripts', 'styles'], function() {
  browserSync.init({
    server: {
      baseDir: "./www"
    }
  });

  gulp.watch('js/*.js', ['scripts']);
  gulp.watch('scss/*.scss', ['styles']);
  gulp.watch('*.html', ['update']);
});

gulp.task('default', ['scripts', 'styles', 'watch']);
