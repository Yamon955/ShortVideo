/**
 * API 接口封装
 * 根据后端 proto 定义的接口进行适配
 */

const API = {
    // 基础配置
    BASE_URL: 'http://localhost:8000', // 网关地址
    
    // 存储 token
    token: localStorage.getItem('token') || '',
    uid: localStorage.getItem('uid') || '',
    
    // 设置 token
    setAuth(token, uid) {
        this.token = token;
        this.uid = uid;
        localStorage.setItem('token', token);
        localStorage.setItem('uid', uid);
    },
    
    // 清除 token
    clearAuth() {
        this.token = '';
        this.uid = '';
        localStorage.removeItem('token');
        localStorage.removeItem('uid');
    },
    
    // 检查是否登录
    isLoggedIn() {
        return !!this.token && !!this.uid;
    },
    
    // 通用请求方法
    async request(url, options = {}) {
        const headers = {
            'Content-Type': 'application/json',
            ...options.headers
        };
        
        // 添加认证头
        if (this.token) {
            headers['Authorization'] = `Bearer ${this.token}`;
        }
        if (this.uid) {
            headers['X-UID'] = this.uid;
        }
        
        try {
            const response = await fetch(`${this.BASE_URL}${url}`, {
                ...options,
                headers
            });
            
            const data = await response.json();
            return data;
        } catch (error) {
            console.error('API 请求失败:', error);
            throw error;
        }
    },
    
    // ========== 用户模块 ==========
    
    /**
     * 用户注册
     * POST /shortvideo/user/register
     * @param {string} username - 用户名
     * @param {string} password - 密码
     */
    async register(username, password) {
        return this.request('/shortvideo/user/register', {
            method: 'POST',
            body: JSON.stringify({
                username,
                password
            })
        });
    },
    
    /**
     * 用户登录
     * POST /shortvideo/user/login
     * @param {string} username - 用户名
     * @param {string} password - 密码
     * @returns {Object} { status_code, status_msg, uid, token }
     */
    async login(username, password) {
        const result = await this.request('/shortvideo/user/login', {
            method: 'POST',
            body: JSON.stringify({
                username,
                password
            })
        });
        
        if (result.status_code === 0 && result.token) {
            this.setAuth(result.token, result.uid);
        }
        
        return result;
    },
    
    /**
     * 退出登录
     */
    logout() {
        this.clearAuth();
    },
    
    /**
     * 批量获取用户资料
     * POST /shortvideo/user
     * @param {number[]} uids - 用户ID列表
     * @param {number[]} profileTypes - 资料类型列表
     *   MAIN_PAGE_INFO = 100 - 个人主页信息
     *   PUBLISH_LIST_COUNT = 101 - 作品列表总数
     *   LIKED_LIST_COUNT = 102 - 赞过列表总数
     *   COLLECT_LIST_COUNT = 103 - 收藏列表总数
     */
    async batchGetProfile(uids, profileTypes = [100, 101, 102, 103]) {
        return this.request('/shortvideo/user', {
            method: 'POST',
            body: JSON.stringify({
                uids,
                profile_types: profileTypes,
                disable_cache: false
            })
        });
    },
    
    /**
     * 设置用户资料
     * POST /shortvideo/user/setprofile
     * @param {Object} profile - 用户资料
     */
    async setProfile(profile) {
        const profileTypes = [];
        if (profile.username !== undefined) profileTypes.push(200);
        if (profile.avator !== undefined) profileTypes.push(201);
        if (profile.sign !== undefined) profileTypes.push(202);
        if (profile.gender !== undefined) profileTypes.push(203);
        
        return this.request('/shortvideo/user/setprofile', {
            method: 'POST',
            body: JSON.stringify({
                uid: parseInt(this.uid),
                profile_types: profileTypes,
                ...profile
            })
        });
    },
    
    // ========== 视频模块 ==========
    
    /**
     * 获取视频流（首页推荐）
     * POST /shortvideo/video/feeds
     * @param {number} startIndex - 起始索引
     * @returns {Object} { next_index, video_infos }
     */
    async getFeeds(startIndex = 0) {
        return this.request('/shortvideo/video/feeds', {
            method: 'POST',
            body: JSON.stringify({
                start_index: startIndex
            })
        });
    },
    
    /**
     * 发布视频
     * POST /shortvideo/video/publish (multipart/form-data)
     * @param {FormData} formData - 包含视频文件和元信息
     */
    async publishVideo(formData) {
        const headers = {};
        if (this.token) {
            headers['Authorization'] = `Bearer ${this.token}`;
        }
        if (this.uid) {
            headers['X-UID'] = this.uid;
        }
        
        try {
            const response = await fetch(`${this.BASE_URL}/shortvideo/video/publish`, {
                method: 'POST',
                headers,
                body: formData
            });
            
            return await response.json();
        } catch (error) {
            console.error('上传视频失败:', error);
            throw error;
        }
    },
    
    /**
     * 获取用户发布的视频列表
     * POST /shortvideo/video/publish-list
     * @param {number} startId - 起始ID
     * @returns {Object} { next_id, is_finash, video_infos }
     */
    async getPublishList(startId = 0) {
        return this.request('/shortvideo/video/publish-list', {
            method: 'POST',
            body: JSON.stringify({
                start_id: startId
            })
        });
    },
    
    /**
     * 获取用户点赞的视频列表
     * POST /shortvideo/video/like-list
     * @param {number} startId - 起始ID
     */
    async getLikeList(startId = 0) {
        return this.request('/shortvideo/video/like-list', {
            method: 'POST',
            body: JSON.stringify({
                start_id: startId
            })
        });
    },
    
    /**
     * 获取用户收藏的视频列表
     * POST /shortvideo/video/collect-list
     * @param {number} startId - 起始ID
     */
    async getCollectList(startId = 0) {
        return this.request('/shortvideo/video/collect-list', {
            method: 'POST',
            body: JSON.stringify({
                start_id: startId
            })
        });
    },
    
    /**
     * 记录观看视频
     * POST /shortvideo/video/watch
     * @param {number} vid - 视频ID
     */
    async watchVideo(vid) {
        return this.request('/shortvideo/video/watch', {
            method: 'POST',
            body: JSON.stringify({
                vid
            })
        });
    },
    
    // ========== 互动模块 ==========
    
    /**
     * 点赞/取消点赞
     * POST /shortvideo/interaction/like
     * @param {number} vid - 视频ID
     * @param {number} operationType - 操作类型 1=点赞 2=取消点赞
     */
    async like(vid, operationType = 1) {
        return this.request('/shortvideo/interaction/like', {
            method: 'POST',
            body: JSON.stringify({
                operation_type: operationType,
                vid
            })
        });
    },
    
    /**
     * 收藏/取消收藏
     * POST /shortvideo/interaction/collect
     * @param {number} vid - 视频ID
     * @param {number} operationType - 操作类型 1=收藏 2=取消收藏
     */
    async collect(vid, operationType = 1) {
        return this.request('/shortvideo/interaction/collect', {
            method: 'POST',
            body: JSON.stringify({
                operation_type: operationType,
                vid
            })
        });
    },
    
    /**
     * 关注/取消关注
     * POST /shortvideo/interaction/follow
     * @param {number} uid - 要关注的用户ID
     * @param {number} operationType - 操作类型 1=关注 2=取消关注
     */
    async follow(uid, operationType = 1) {
        return this.request('/shortvideo/interaction/follow', {
            method: 'POST',
            body: JSON.stringify({
                operation_type: operationType,
                uid
            })
        });
    },
    
    /**
     * 评论/撤回评论
     * POST /shortvideo/interaction/comment
     * @param {number} vid - 视频ID
     * @param {string} msg - 评论内容
     * @param {number} operationType - 操作类型 1=评论 2=撤回评论
     */
    async comment(vid, msg, operationType = 1) {
        return this.request('/shortvideo/interaction/comment', {
            method: 'POST',
            body: JSON.stringify({
                operation_type: operationType,
                vid,
                msg
            })
        });
    },
    
    /**
     * 举报视频
     * POST /shortvideo/interaction/report
     * @param {number} reportVid - 被举报视频ID
     * @param {string} reportUid - 被举报用户ID
     * @param {number[]} reportTypes - 举报类型列表
     * @param {string} extInfo - 补充说明
     */
    async report(reportVid, reportUid, reportTypes, extInfo = '') {
        return this.request('/shortvideo/interaction/report', {
            method: 'POST',
            body: JSON.stringify({
                operation_type: 1,
                uid: parseInt(this.uid),
                report_vid: reportVid,
                report_uid: reportUid,
                report_types: reportTypes,
                ext_info: extInfo
            })
        });
    },
    
    // ========== 推荐模块 ==========
    
    /**
     * 获取推荐视频列表
     * POST /shortvideo/recommend/feeds
     * @param {number} startIndex - 起始索引
     */
    async getFeedsRecommend(startIndex = 0) {
        return this.request('/shortvideo/recommend/feeds', {
            method: 'POST',
            body: JSON.stringify({
                uid: parseInt(this.uid) || 0,
                start_index: startIndex
            })
        });
    }
};

// 视频分类标签映射
const VIDEO_TAGS = {
    0: '娱乐',
    1: '商业',
    2: '音乐',
    3: '生活',
    4: '教育',
    5: '科技',
    6: '时尚',
    7: '运动',
    8: '汽车',
    9: '艺术',
    10: '新闻',
    11: '游戏',
    12: '健康',
    13: 'DIY',
    14: '文化'
};

// 举报类型映射
const REPORT_TYPES = {
    1: '暴力行为',
    2: '恐怖主义',
    3: '歧视性内容',
    4: '色情内容',
    5: '虚假信息',
    6: '版权侵权',
    7: '未成年安全',
    8: '虐待动物',
    9: '侵犯隐私',
    10: '引人不适'
};

// 导出
window.API = API;
window.VIDEO_TAGS = VIDEO_TAGS;
window.REPORT_TYPES = REPORT_TYPES;
