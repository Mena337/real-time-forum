function showPage(pageId) {
    const pages = document.querySelectorAll('.page');
    pages.forEach(page => page.classList.remove('active'));

    const page = document.getElementById(pageId);
    page.classList.add('active');
}

window.onload = function() {
    showPage('login_page'); 
};
