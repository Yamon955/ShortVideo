/**
 * 短视频网站主应用逻辑
 */

class ShortVideoApp {
    constructor() {
        this.currentPage = 'home';
        this.currentVideoIndex = 0;
        this.videoList = [];
        this.isAuthMode = 'login';
        this.currentVideoDetail = null;
        this.userProfile = null;
        this.feedsNextIndex = 0;
        this.init();
    }
    
    init() {
        this.bindEvents();
        this.checkAuthStatus();
        this.loadFeeds();
    }
    
    bindEvents() {
        // 导航菜单
        document.querySelectorAll('.nav-item').forEach(item => {
            item.addEventListener('click', () => this.switchPage(item.dataset.page));
        });
        
        // 登录相关
        document.getElementById('btnShowLogin').addEventListener('click', () => this.showLoginModal());
        document.getElementById('btnCloseLogin').addEventListener('click', () => this.hideLoginModal());
        document.getElementById('authSwitchLink').addEventListener('click', (e) => {
            e.preventDefault();
            this.toggleAuthMode();
        });
        document.getElementById('authForm').addEventListener('submit', (e) => this.handleAuth(e));
        
        // 编辑资料
        document.getElementById('btnEditProfile').addEventListener('click', () => this.showEditProfileModal());
        document.getElementById('btnCloseEditProfile').addEventListener('click', () => this.hideEditProfileModal());
        document.getElementById('editProfileForm').addEventListener('submit', (e) => this.handleEditProfile(e));
        
        // 视频上传
        const uploadArea = document.getElementById('uploadArea');
        uploadArea.addEventListener('click', () => document.getElementById('videoFile').click());
        document.getElementById('videoFile').addEventListener('change', (e) => this.handleVideoSelect(e));
        document.getElementById('btnRemoveVideo').addEventListener('click', () => this.removeSelectedVideo());
        
        // 封面上传
        document.getElementById('coverUploadArea').addEventListener('click', () => document.getElementById('coverFile').click());
        document.getElementById('coverFile').addEventListener('change', (e) => this.handleCoverSelect(e));
        document.getElementById('btnRemoveCover').addEventListener('click', () => this.removeSelectedCover());
        
        // 标题字数
        document.getElementById('videoTitle').addEventListener('input', (e) => {
            document.getElementById('titleCount').textContent = e.target.value.length;
        });
        
        // 发布视频
        document.getElementById('uploadForm').addEventListener('submit', (e) => this.handlePublishVideo(e));
        
        // 视频滚动
        document.getElementById('btnPrevVideo').addEventListener('click', () => this.prevVideo());
        document.getElementById('btnNextVideo').addEventListener('click', () => this.nextVideo());
        
        document.addEventListener('keydown', (e) => {
            if (this.currentPage === 'home') {
                if (e.key === 'ArrowUp') this.prevVideo();
                if (e.key === 'ArrowDown') this.nextVideo();
            }
        });
        
        document.getElementById('videoContainer').addEventListener('wheel', (e) => {
            e.preventDefault();
            if (e.deltaY > 0) this.nextVideo();
            else this.prevVideo();
        });
        
        // 分类标签
        document.querySelectorAll('.tags-container .tag-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                document.querySelectorAll('.tags-container .tag-btn').forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                this.loadDiscoverVideos(btn.dataset.tag);
            });
        });
        
        // 个人中心标签
        document.querySelectorAll('.profile-tabs .tab-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                document.querySelectorAll('.profile-tabs .tab-btn').forEach(b => b.classList.remove('active'));
                document.querySelectorAll('.tab-content').forEach(c => c.classList.remove('active'));
                btn.classList.add('active');
                document.getElementById(`${btn.dataset.tab}Content`).classList.add('active');
                this.loadTabContent(btn.dataset.tab);
            });
        });
        
        // 视频详情
        document.getElementById('btnCloseVideoDetail').addEventListener('click', () => this.hideVideoDetailModal());
        document.getElementById('btnLike').addEventListener('click', () => this.handleLike());
        document.getElementById('btnCollect').addEventListener('click', () => this.handleCollect());
        document.getElementById('btnFollow').addEventListener('click', () => this.handleFollow());
        document.getElementById('btnSendComment').addEventListener('click', () => this.handleSendComment());
        document.getElementById('btnReport').addEventListener('click', () => this.showReportModal());
        
        // 举报
        document.getElementById('btnCloseReport').addEventListener('click', () => this.hideReportModal());
        document.getElementById('reportForm').addEventListener('submit', (e) => this.handleReport(e));
        
        document.getElementById('commentInput').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') this.handleSendComment();
        });
    }
    
    switchPage(page) {
        this.currentPage = page;
        document.querySelectorAll('.nav-item').forEach(item => {
            item.classList.toggle('active', item.dataset.page === page);
        });
        document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
        document.getElementById(`${page}Page`).classList.add('active');
        
        switch (page) {
            case 'home':
                if (this.videoList.length === 0) this.loadFeeds();
                break;
            case 'discover':
                this.loadDiscoverVideos();
                break;
            case 'upload':
            case 'profile':
                if (!API.isLoggedIn()) {
                    this.showToast('请先登录', 'error');
                    this.showLoginModal();
                    this.switchPage('home');
                } else if (page === 'profile') {
                    this.loadProfile();
                }
                break;
        }
    }
    
    checkAuthStatus() {
        if (API.isLoggedIn()) {
            this.showLoggedInState();
            this.loadUserProfile();
        }
    }
    
    showLoginModal() {
        document.getElementById('loginModal').classList.add('show');
    }
    
    hideLoginModal() {
        document.getElementById('loginModal').classList.remove('show');
        document.getElementById('authForm').reset();
    }
    
    toggleAuthMode() {
        this.isAuthMode = this.isAuthMode === 'login' ? 'register' : 'login';
        const isLogin = this.isAuthMode === 'login';
        document.getElementById('authTitle').textContent = isLogin ? '登录' : '注册';
        document.getElementById('btnSubmitAuth').textContent = isLogin ? '登录' : '注册';
        document.getElementById('confirmPasswordGroup').style.display = isLogin ? 'none' : 'block';
        document.getElementById('authSwitchText').textContent = isLogin ? '还没有账号？' : '已有账号？';
        document.getElementById('authSwitchLink').textContent = isLogin ? '立即注册' : '立即登录';
    }
    
    async handleAuth(e) {
        e.preventDefault();
        const username = document.getElementById('authUsername').value.trim();
        const password = document.getElementById('authPassword').value;
        
        if (!username || !password) {
            this.showToast('请填写完整信息', 'error');
            return;
        }
        
        this.showLoading();
        try {
            if (this.isAuthMode === 'login') {
                const result = await API.login(username, password);
                if (result.status_code === 0) {
                    this.showToast('登录成功');
                    this.hideLoginModal();
                    this.showLoggedInState();
                    this.loadUserProfile();
                } else {
                    this.showToast(result.status_msg || '登录失败', 'error');
                }
            } else {
                const confirmPassword = document.getElementById('authConfirmPassword').value;
                if (password !== confirmPassword) {
                    this.showToast('两次密码不一致', 'error');
                    this.hideLoading();
                    return;
                }
                const result = await API.register(username, password);
                if (result.status_code === 0) {
                    this.showToast('注册成功，请登录');
                    this.toggleAuthMode();
                } else {
                    this.showToast(result.status_msg || '注册失败', 'error');
                }
            }
        } catch (error) {
            this.showToast('网络错误', 'error');
        }
        this.hideLoading();
    }
    
    showLoggedInState() {
        const userSection = document.getElementById('userSection');
        userSection.innerHTML = `
            <div class="user-logged" onclick="app.switchPage('profile')">
<img src="static/images/default-avatar.svg" alt="头像" id="sidebarAvatar">
                <div class="user-info">
                    <span class="username" id="sidebarUsername">用户</span>
                </div>
                <button class="logout-btn" onclick="event.stopPropagation(); app.handleLogout()">
                    <i class="fas fa-sign-out-alt"></i>
                </button>
            </div>
        `;
    }
    
    handleLogout() {
        API.logout();
        location.reload();
    }
    
    async loadUserProfile() {
        if (!API.isLoggedIn()) return;
        try {
            const result = await API.batchGetProfile([parseInt(API.uid)]);
            if (result.user_infos && result.user_infos.length > 0) {
                this.userProfile = result.user_infos[0];
                this.updateUserDisplay();
            }
        } catch (error) {
            console.error('获取用户资料失败:', error);
        }
    }
    
    updateUserDisplay() {
        if (!this.userProfile) return;
        const mainInfo = this.userProfile.main_page_info || {};
        
        const sidebarUsername = document.getElementById('sidebarUsername');
        const sidebarAvatar = document.getElementById('sidebarAvatar');
        if (sidebarUsername) sidebarUsername.textContent = mainInfo.username || '用户';
        if (sidebarAvatar && mainInfo.avator) sidebarAvatar.src = mainInfo.avator;
        
        document.getElementById('userName').textContent = mainInfo.username || '用户';
        document.getElementById('userSign').textContent = mainInfo.sign || '这个人很懒，什么都没写';
        if (mainInfo.avator) document.getElementById('userAvatar').src = mainInfo.avator;
        document.getElementById('fansCount').textContent = mainInfo.fans_count || 0;
        document.getElementById('followsCount').textContent = mainInfo.follows_count || 0;
        document.getElementById('publishCount').textContent = this.userProfile.publish_list_count || 0;
        document.getElementById('likedCount').textContent = this.userProfile.liked_list_count || 0;
        document.getElementById('worksTabCount').textContent = this.userProfile.publish_list_count || 0;
        document.getElementById('likedTabCount').textContent = this.userProfile.liked_list_count || 0;
        document.getElementById('collectedTabCount').textContent = this.userProfile.collect_list_count || 0;
    }
    
    showEditProfileModal() {
        if (!this.userProfile) return;
        const mainInfo = this.userProfile.main_page_info || {};
        document.getElementById('editUsername').value = mainInfo.username || '';
        document.getElementById('editSign').value = mainInfo.sign || '';
        document.querySelectorAll('input[name="gender"]').forEach(input => {
            input.checked = parseInt(input.value) === (mainInfo.gender || 0);
        });
        document.getElementById('editProfileModal').classList.add('show');
    }
    
    hideEditProfileModal() {
        document.getElementById('editProfileModal').classList.remove('show');
    }
    
    async handleEditProfile(e) {
        e.preventDefault();
        const username = document.getElementById('editUsername').value.trim();
        const sign = document.getElementById('editSign').value.trim();
        const gender = parseInt(document.querySelector('input[name="gender"]:checked').value);
        
        this.showLoading();
        try {
            const result = await API.setProfile({ username, sign, gender });
            if (result.status_code === 0) {
                this.showToast('资料更新成功');
                this.hideEditProfileModal();
                this.loadUserProfile();
            } else {
                this.showToast(result.status_msg || '更新失败', 'error');
            }
        } catch (error) {
            this.showToast('网络错误', 'error');
        }
        this.hideLoading();
    }
    
    async loadFeeds() {
        this.showLoading();
        try {
            const result = await API.getFeeds(this.feedsNextIndex);
            if (result.video_infos && result.video_infos.length > 0) {
                this.videoList = [...this.videoList, ...result.video_infos];
                this.feedsNextIndex = result.next_index || 0;
                this.renderVideoSlides();
            } else if (this.videoList.length === 0) {
                this.renderDemoVideos();
            }
        } catch (error) {
            if (this.videoList.length === 0) this.renderDemoVideos();
        }
        this.hideLoading();
    }
    
    renderVideoSlides() {
        const container = document.getElementById('videoContainer');
        container.innerHTML = '';
        
        this.videoList.forEach((video, index) => {
            const slide = document.createElement('div');
            slide.className = 'video-slide';
            slide.style.transform = `translateY(${(index - this.currentVideoIndex) * 100}%)`;
            slide.innerHTML = `
                <video src="${video.video_url}" loop playsinline></video>
                <div class="video-info-overlay">
                    <div class="author">
<img src="${video.author_avatar || 'static/images/default-avatar.svg'}" alt="作者">
                        <span class="name">@${video.author_name || '用户'}</span>
                    </div>
                    <p class="title">${video.title || ''}</p>
                </div>
                <div class="video-actions">
                    <button class="video-action-btn ${video.liked ? 'liked' : ''}" onclick="app.quickLike(${video.vid})">
                        <i class="${video.liked ? 'fas' : 'far'} fa-heart"></i>
                        <span>${this.formatNumber(video.likes_count || 0)}</span>
                    </button>
                    <button class="video-action-btn ${video.collected ? 'collected' : ''}" onclick="app.quickCollect(${video.vid})">
                        <i class="${video.collected ? 'fas' : 'far'} fa-star"></i>
                        <span>${this.formatNumber(video.collect_count || 0)}</span>
                    </button>
                    <button class="video-action-btn" onclick="app.showVideoDetail(${index})">
                        <i class="far fa-comment"></i>
                        <span>${this.formatNumber(video.comments_count || 0)}</span>
                    </button>
                    <button class="video-action-btn" onclick="app.shareVideo(${video.vid})">
                        <i class="fas fa-share"></i>
                        <span>分享</span>
                    </button>
                </div>
            `;
            container.appendChild(slide);
        });
        this.playCurrentVideo();
    }
    
    renderDemoVideos() {
        this.videoList = [
            { vid: 1, video_url: 'https://www.w3schools.com/html/mov_bbb.mp4', title: '欢迎来到短视频平台！', likes_count: 1234, collect_count: 567, comments_count: 89, liked: false, collected: false, author_name: '官方账号' },
            { vid: 2, video_url: 'https://www.w3schools.com/html/movie.mp4', title: '发现更多精彩内容！', likes_count: 2345, collect_count: 678, comments_count: 123, liked: false, collected: false, author_name: '创作者' }
        ];
        this.renderVideoSlides();
    }
    
    prevVideo() {
        if (this.currentVideoIndex > 0) {
            this.currentVideoIndex--;
            this.updateVideoPosition();
        }
    }
    
    nextVideo() {
        if (this.currentVideoIndex < this.videoList.length - 1) {
            this.currentVideoIndex++;
            this.updateVideoPosition();
            if (this.currentVideoIndex >= this.videoList.length - 2) this.loadFeeds();
        }
    }
    
    updateVideoPosition() {
        document.querySelectorAll('.video-slide').forEach((slide, index) => {
            slide.style.transform = `translateY(${(index - this.currentVideoIndex) * 100}%)`;
        });
        this.playCurrentVideo();
    }
    
    playCurrentVideo() {
        document.querySelectorAll('.video-slide video').forEach((video, index) => {
            if (index === this.currentVideoIndex) video.play().catch(() => {});
            else video.pause();
        });
        if (this.videoList[this.currentVideoIndex]) {
            API.watchVideo(this.videoList[this.currentVideoIndex].vid).catch(() => {});
        }
    }
    
    async quickLike(vid) {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        const video = this.videoList.find(v => v.vid === vid);
        if (!video) return;
        try {
            const result = await API.like(vid, video.liked ? 2 : 1);
            if (result.status_code === 0) {
                video.liked = !video.liked;
                video.likes_count += video.liked ? 1 : -1;
                this.renderVideoSlides();
                this.showToast(video.liked ? '已点赞' : '已取消点赞');
            }
        } catch (error) { this.showToast('操作失败', 'error'); }
    }
    
    async quickCollect(vid) {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        const video = this.videoList.find(v => v.vid === vid);
        if (!video) return;
        try {
            const result = await API.collect(vid, video.collected ? 2 : 1);
            if (result.status_code === 0) {
                video.collected = !video.collected;
                video.collect_count += video.collected ? 1 : -1;
                this.renderVideoSlides();
                this.showToast(video.collected ? '已收藏' : '已取消收藏');
            }
        } catch (error) { this.showToast('操作失败', 'error'); }
    }
    
    shareVideo(vid) {
        const url = `${location.origin}?v=${vid}`;
        if (navigator.clipboard) { navigator.clipboard.writeText(url); this.showToast('链接已复制'); }
        else this.showToast('分享链接: ' + url);
    }
    
    async loadDiscoverVideos(tag = null) {
        const grid = document.getElementById('discoverGrid');
        grid.innerHTML = '<div class="loading-placeholder">加载中...</div>';
        try {
            const result = await API.getFeeds(0);
            let videos = result.video_infos || [];
            if (tag !== null) videos = videos.filter(v => v.tags == tag);
            this.renderDiscoverGrid(videos.length ? videos : this.videoList);
        } catch (error) { this.renderDiscoverGrid(this.videoList); }
    }
    
    renderDiscoverGrid(videos) {
        const grid = document.getElementById('discoverGrid');
        if (!videos.length) { grid.innerHTML = '<div class="empty-state"><i class="fas fa-search"></i><p>暂无相关视频</p></div>'; return; }
        grid.innerHTML = videos.map((video, i) => `
            <div class="video-card" onclick="app.openVideoDetail(${JSON.stringify(video).replace(/"/g, '&quot;')})">
<div class="thumbnail"><img src="${video.cover_url || 'static/images/default-cover.svg'}" alt="${video.title}"><i class="fas fa-play-circle play-icon"></i></div>
                <div class="card-info"><p class="card-title">${video.title || '无标题'}</p><div class="card-stats"><span><i class="fas fa-heart"></i> ${this.formatNumber(video.likes_count || 0)}</span><span><i class="fas fa-comment"></i> ${this.formatNumber(video.comments_count || 0)}</span></div></div>
            </div>
        `).join('');
    }
    
    handleVideoSelect(e) { const file = e.target.files[0]; if (file) this.selectVideo(file); }
    
    selectVideo(file) {
        this.selectedVideoFile = file;
        document.getElementById('uploadArea').style.display = 'none';
        document.getElementById('previewArea').style.display = 'block';
        document.getElementById('videoPreview').src = URL.createObjectURL(file);
    }
    
    removeSelectedVideo() {
        this.selectedVideoFile = null;
        document.getElementById('uploadArea').style.display = 'block';
        document.getElementById('previewArea').style.display = 'none';
        document.getElementById('videoPreview').src = '';
        document.getElementById('videoFile').value = '';
    }
    
    handleCoverSelect(e) {
        const file = e.target.files[0];
        if (file) {
            this.selectedCoverFile = file;
            document.getElementById('coverUploadArea').style.display = 'none';
            document.getElementById('coverPreview').style.display = 'block';
            const reader = new FileReader();
            reader.onload = (e) => document.getElementById('coverImg').src = e.target.result;
            reader.readAsDataURL(file);
        }
    }
    
    removeSelectedCover() {
        this.selectedCoverFile = null;
        document.getElementById('coverUploadArea').style.display = 'block';
        document.getElementById('coverPreview').style.display = 'none';
        document.getElementById('coverImg').src = '';
        document.getElementById('coverFile').value = '';
    }
    
    async handlePublishVideo(e) {
        e.preventDefault();
        if (!this.selectedVideoFile) { this.showToast('请选择视频文件', 'error'); return; }
        const title = document.getElementById('videoTitle').value.trim();
        if (!title) { this.showToast('请输入视频标题', 'error'); return; }
        
        const formData = new FormData();
        formData.append('video', this.selectedVideoFile);
        formData.append('title', title);
        formData.append('tags', document.querySelector('input[name="tags"]:checked').value);
        if (this.selectedCoverFile) formData.append('cover', this.selectedCoverFile);
        
        this.showLoading();
        try {
            const result = await API.publishVideo(formData);
            if (result.status_code === 0) { this.showToast('发布成功！'); this.resetUploadForm(); this.switchPage('profile'); }
            else this.showToast(result.status_msg || '发布失败', 'error');
        } catch (error) { this.showToast('上传失败', 'error'); }
        this.hideLoading();
    }
    
    resetUploadForm() { this.removeSelectedVideo(); this.removeSelectedCover(); document.getElementById('uploadForm').reset(); document.getElementById('titleCount').textContent = '0'; }
    
    async loadProfile() { await this.loadUserProfile(); this.loadTabContent('works'); }
    
    async loadTabContent(tab) {
        const gridId = { works: 'worksGrid', liked: 'likedGrid', collected: 'collectedGrid' }[tab];
        const grid = document.getElementById(gridId);
        try {
            let result;
            if (tab === 'works') result = await API.getPublishList(0);
            else if (tab === 'liked') result = await API.getLikeList(0);
            else result = await API.getCollectList(0);
            
            if (result.video_infos?.length) this.renderProfileGrid(grid, result.video_infos);
            else grid.innerHTML = `<div class="empty-state"><i class="fas fa-inbox"></i><p>暂无内容</p></div>`;
        } catch (error) { grid.innerHTML = '<div class="empty-state"><i class="fas fa-inbox"></i><p>加载失败</p></div>'; }
    }
    
    renderProfileGrid(grid, videos) {
        grid.innerHTML = videos.map(video => `
            <div class="video-card" onclick="app.openVideoDetail(${JSON.stringify(video).replace(/"/g, '&quot;')})">
<div class="thumbnail"><img src="${video.cover_url || 'static/images/default-cover.svg'}" alt="${video.title}"><i class="fas fa-play-circle play-icon"></i></div>
                <div class="card-info"><p class="card-title">${video.title || '无标题'}</p><div class="card-stats"><span><i class="fas fa-heart"></i> ${this.formatNumber(video.likes_count || 0)}</span></div></div>
            </div>
        `).join('');
    }
    
    openVideoDetail(video) { this.currentVideoDetail = video; this.showVideoDetailModal(); }
    showVideoDetail(index) { this.currentVideoDetail = this.videoList[index]; this.showVideoDetailModal(); }
    
    showVideoDetailModal() {
        if (!this.currentVideoDetail) return;
        const video = this.currentVideoDetail;
        document.getElementById('detailVideo').src = video.video_url;
document.getElementById('authorAvatar').src = video.author_avatar || 'static/images/default-avatar.svg';
        document.getElementById('authorName').textContent = '@' + (video.author_name || '用户');
        document.getElementById('publishTime').textContent = this.formatTime(video.publish_time);
        document.getElementById('videoTitle2').textContent = video.title || '';
        document.getElementById('likeCount').textContent = this.formatNumber(video.likes_count || 0);
        document.getElementById('collectCount').textContent = this.formatNumber(video.collect_count || 0);
        document.getElementById('commentCount').textContent = this.formatNumber(video.comments_count || 0);
        
        const likeBtn = document.getElementById('btnLike');
        likeBtn.classList.toggle('active', video.liked);
        likeBtn.querySelector('i').className = video.liked ? 'fas fa-heart' : 'far fa-heart';
        
        const collectBtn = document.getElementById('btnCollect');
        collectBtn.classList.toggle('active', video.collected);
        collectBtn.querySelector('i').className = video.collected ? 'fas fa-star' : 'far fa-star';
        
        document.getElementById('commentsList').innerHTML = '<div class="empty-state"><p>暂无评论</p></div>';
        document.getElementById('videoDetailModal').classList.add('show');
    }
    
    hideVideoDetailModal() { document.getElementById('videoDetailModal').classList.remove('show'); document.getElementById('detailVideo').pause(); this.currentVideoDetail = null; }
    
    async handleLike() {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        if (!this.currentVideoDetail) return;
        const video = this.currentVideoDetail;
        try {
            const result = await API.like(video.vid, video.liked ? 2 : 1);
            if (result.status_code === 0) {
                video.liked = !video.liked;
                video.likes_count += video.liked ? 1 : -1;
                const likeBtn = document.getElementById('btnLike');
                likeBtn.classList.toggle('active', video.liked);
                likeBtn.querySelector('i').className = video.liked ? 'fas fa-heart' : 'far fa-heart';
                document.getElementById('likeCount').textContent = this.formatNumber(video.likes_count);
                this.showToast(video.liked ? '已点赞' : '已取消点赞');
            }
        } catch (error) { this.showToast('操作失败', 'error'); }
    }
    
    async handleCollect() {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        if (!this.currentVideoDetail) return;
        const video = this.currentVideoDetail;
        try {
            const result = await API.collect(video.vid, video.collected ? 2 : 1);
            if (result.status_code === 0) {
                video.collected = !video.collected;
                video.collect_count += video.collected ? 1 : -1;
                const collectBtn = document.getElementById('btnCollect');
                collectBtn.classList.toggle('active', video.collected);
                collectBtn.querySelector('i').className = video.collected ? 'fas fa-star' : 'far fa-star';
                document.getElementById('collectCount').textContent = this.formatNumber(video.collect_count);
                this.showToast(video.collected ? '已收藏' : '已取消收藏');
            }
        } catch (error) { this.showToast('操作失败', 'error'); }
    }
    
    async handleFollow() {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        if (!this.currentVideoDetail?.author_uid) { this.showToast('无法获取作者信息', 'error'); return; }
        const followBtn = document.getElementById('btnFollow');
        const isFollowing = followBtn.classList.contains('following');
        try {
            const result = await API.follow(this.currentVideoDetail.author_uid, isFollowing ? 2 : 1);
            if (result.status_code === 0) {
                followBtn.classList.toggle('following');
                followBtn.innerHTML = isFollowing ? '<i class="fas fa-plus"></i> 关注' : '<i class="fas fa-check"></i> 已关注';
                this.showToast(isFollowing ? '已取消关注' : '关注成功');
            }
        } catch (error) { this.showToast('操作失败', 'error'); }
    }
    
    async handleSendComment() {
        if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; }
        if (!this.currentVideoDetail) return;
        const input = document.getElementById('commentInput');
        const content = input.value.trim();
        if (!content) { this.showToast('请输入评论内容', 'error'); return; }
        
        try {
            const result = await API.comment(this.currentVideoDetail.vid, content, 1);
            if (result.status_code === 0) {
                input.value = '';
                this.currentVideoDetail.comments_count++;
                document.getElementById('commentCount').textContent = this.formatNumber(this.currentVideoDetail.comments_count);
                const commentsList = document.getElementById('commentsList');
                const emptyState = commentsList.querySelector('.empty-state');
                if (emptyState) emptyState.remove();
                const commentItem = document.createElement('div');
                commentItem.className = 'comment-item';
                commentItem.innerHTML = `<img src="${this.userProfile?.main_page_info?.avator || 'static/images/default-avatar.svg'}" alt="头像"><div class="comment-content"><span class="comment-author">${this.userProfile?.main_page_info?.username || '我'}</span><p class="comment-text">${content}</p></div>`;
                commentsList.insertBefore(commentItem, commentsList.firstChild);
                this.showToast('评论成功');
            }
        } catch (error) { this.showToast('评论失败', 'error'); }
    }
    
    showReportModal() { if (!API.isLoggedIn()) { this.showToast('请先登录', 'error'); this.showLoginModal(); return; } document.getElementById('reportModal').classList.add('show'); }
    hideReportModal() { document.getElementById('reportModal').classList.remove('show'); document.getElementById('reportForm').reset(); }
    
    async handleReport(e) {
        e.preventDefault();
        if (!this.currentVideoDetail) return;
        const reportTypes = Array.from(document.querySelectorAll('input[name="reportType"]:checked')).map(input => parseInt(input.value));
        if (!reportTypes.length) { this.showToast('请选择举报原因', 'error'); return; }
        
        this.showLoading();
        try {
            const result = await API.report(this.currentVideoDetail.vid, this.currentVideoDetail.author_uid || '', reportTypes, document.getElementById('reportDescription').value.trim());
            if (result.status_code === 0) { this.showToast('举报已提交'); this.hideReportModal(); }
            else this.showToast(result.status_msg || '举报失败', 'error');
        } catch (error) { this.showToast('提交失败', 'error'); }
        this.hideLoading();
    }
    
    formatNumber(num) { if (num >= 10000) return (num / 10000).toFixed(1) + 'w'; if (num >= 1000) return (num / 1000).toFixed(1) + 'k'; return num.toString(); }
    formatTime(timestamp) { if (!timestamp) return ''; const d = new Date(timestamp * 1000), now = new Date(), diff = now - d; if (diff < 60000) return '刚刚'; if (diff < 3600000) return Math.floor(diff / 60000) + '分钟前'; if (diff < 86400000) return Math.floor(diff / 3600000) + '小时前'; return `${d.getMonth() + 1}月${d.getDate()}日`; }
    
    showToast(message, type = 'success') {
        const toast = document.getElementById('toast');
        toast.className = `toast ${type}`;
        document.getElementById('toastMessage').textContent = message;
        toast.querySelector('i').className = type === 'success' ? 'fas fa-check-circle' : 'fas fa-exclamation-circle';
        toast.classList.add('show');
        setTimeout(() => toast.classList.remove('show'), 3000);
    }
    
    showLoading() { document.getElementById('loading').classList.add('show'); }
    hideLoading() { document.getElementById('loading').classList.remove('show'); }
}

let app;
document.addEventListener('DOMContentLoaded', () => { app = new ShortVideoApp(); });