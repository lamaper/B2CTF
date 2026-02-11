# 前端迁移指南 - B2CTF 权限系统更新

## 📋 概述

由于后端权限系统的升级，前端需要进行以下调整。本指南提供了具体的代码修改方案。

**影响范围**: 3 个文件，3 个 API endpoint  
**预计工作量**: 30-45 分钟  
**难度等级**: ⭐⭐ 简单

---

## 🔴 必须修改 (Breaking Changes)

### 1. 创建题目页面

**文件**: `frontend/src/views/CreateChallenge.vue`

#### 问题
端点已从 `/challenge` 迁移到 `/admin/challenge`

#### 修改方案

**查找现有代码**:
```javascript
// 在 CreateChallenge.vue 中查找类似代码
http.post('/challenge', {...})

// 或在 src/api 中查找
import { post } from '@/api/http'
post('/challenge', {...})
```

**修改为**:
```javascript
http.post('/admin/challenge', {...})

// 或
post('/admin/challenge', {...})
```

#### 完整示例

<details>
<summary>点击展开完整代码示例</summary>

```vue
<template>
  <div class="create-challenge">
    <form @submit="handleSubmit">
      <!-- 表单字段 -->
      <input v-model="form.title" placeholder="题目标题">
      <input v-model="form.category" placeholder="分类">
      <input v-model="form.description" placeholder="描述">
      <input v-model="form.flag" placeholder="flag">
      <input v-model="form.score" type="number" placeholder="分数">
      <input v-model="form.competition_id" type="number" placeholder="比赛 ID">
      
      <!-- 新增：动态题参数（可选） -->
      <div v-if="form.is_dynamic">
        <input v-model="form.image_name" placeholder="Docker 镜像名称">
        <input v-model="form.container_port" type="number" placeholder="容器端口">
        <input v-model="form.memory_limit" placeholder="内存限制，如 256m">
        <input v-model="form.cpu_limit" placeholder="CPU 限制，如 0.5">
      </div>
      
      <button type="submit">创建题目</button>
    </form>
  </div>
</template>

<script>
import { post } from '@/api/http'

export default {
  name: 'CreateChallenge',
  data() {
    return {
      form: {
        title: '',
        category: '',
        description: '',
        flag: '',
        score: 0,
        competition_id: 0,
        tags: [],
        is_dynamic: false,
        image_name: '',
        container_port: 0,
        memory_limit: '',
        cpu_limit: ''
      }
    }
  },
  methods: {
    async handleSubmit() {
      try {
        // ✅ 改为 /admin/challenge
        const response = await post('/admin/challenge', this.form)
        
        if (response.code === 200) {
          alert('题目创建成功!')
          this.$router.push('/challenges')
        } else {
          alert(`创建失败: ${response.msg}`)
        }
      } catch (error) {
        console.error('创建题目失败:', error)
        alert('创建题目出错')
      }
    }
  }
}
</script>
```

</details>

---

### 2. 创建比赛页面

**文件**: `frontend/src/views/CreateCompetition.vue`

#### 问题
1. 端点已从 `/competitions` 迁移到 `/admin/competition`
2. 新增强制参数 `mode` (个人赛=0，团队赛=1)

#### 修改方案

**第 1 步：更改 endpoint**

```javascript
// 改前
await post('/competitions', competitionData)

// 改后
await post('/admin/competition', competitionData)
```

**第 2 步：添加 `mode` 字段**

```javascript
// 改前
const competitionData = {
  title: this.form.title,
  description: this.form.description,
  type: this.form.type,
  start_time: this.form.startTime,
  end_time: this.form.endTime
}

// 改后
const competitionData = {
  title: this.form.title,
  description: this.form.description,
  type: this.form.type,
  mode: this.form.mode,  // ⭐ 新增（0=个人赛，1=团队赛）
  start_time: this.form.startTime,
  end_time: this.form.endTime
}
```

#### 完整示例

<details>
<summary>点击展开完整代码示例</summary>

```vue
<template>
  <div class="create-competition">
    <form @submit="handleSubmit">
      <div>
        <label for="title">比赛标题</label>
        <input id="title" v-model="form.title" required>
      </div>
      
      <div>
        <label for="description">比赛描述</label>
        <textarea id="description" v-model="form.description"></textarea>
      </div>
      
      <div>
        <label for="type">比赛类型</label>
        <select id="type" v-model="form.type">
          <option value="0">标准 CTF</option>
          <option value="1">靶场</option>
          <option value="2">练习赛</option>
        </select>
      </div>
      
      <!-- ⭐ 新增：选择比赛模式 -->
      <div>
        <label for="mode">比赛模式</label>
        <select id="mode" v-model="form.mode">
          <option value="0">个人赛</option>
          <option value="1">团队赛</option>
        </select>
      </div>
      
      <div>
        <label for="startTime">开始时间</label>
        <input id="startTime" v-model="form.startTime" type="datetime-local" required>
      </div>
      
      <div>
        <label for="endTime">结束时间</label>
        <input id="endTime" v-model="form.endTime" type="datetime-local" required>
      </div>
      
      <button type="submit">创建比赛</button>
    </form>
  </div>
</template>

<script>
import { post } from '@/api/http'

export default {
  name: 'CreateCompetition',
  data() {
    return {
      form: {
        title: '',
        description: '',
        type: '0',
        mode: '0',  // ⭐ 新增默认值
        startTime: '',
        endTime: ''
      }
    }
  },
  methods: {
    async handleSubmit() {
      try {
        // ✅ 检查开始时间 < 结束时间
        const start = new Date(this.form.startTime)
        const end = new Date(this.form.endTime)
        
        if (start >= end) {
          alert('开始时间必须早于结束时间')
          return
        }
        
        const competitionData = {
          title: this.form.title,
          description: this.form.description,
          type: parseInt(this.form.type),
          mode: parseInt(this.form.mode),  // ⭐ 新增
          start_time: start.toISOString(),  // 转换为 ISO 格式
          end_time: end.toISOString()       // 转换为 ISO 格式
        }
        
        // ✅ 改为 /admin/competition
        const response = await post('/admin/competition', competitionData)
        
        if (response.code === 200) {
          alert('比赛创建成功!')
          this.$router.push('/competitions')
        } else {
          alert(`创建失败: ${response.msg}`)
        }
      } catch (error) {
        console.error('创建比赛失败:', error)
        alert('创建比赛出错')
      }
    }
  }
}
</script>

<style scoped>
input, textarea, select {
  display: block;
  margin: 10px 0;
  padding: 8px;
  width: 100%;
  max-width: 400px;
}

button {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 20px;
}

button:hover {
  background-color: #45a049;
}
</style>
```

</details>

---

### 3. 文件上传功能

**文件**: 根据项目结构可能在以下位置：
- `frontend/src/api/*.js`
- `frontend/src/views/CreateChallenge.vue`
- `frontend/src/components/Upload.vue`

#### 问题
端点已从 `/upload` 迁移到 `/admin/upload`

#### 修改方案

**查找现有代码**:
```javascript
// 在上传文件的代码中查找
formData.post('/upload')
// 或
axios.post('/upload', formData)
// 或
fetch('/upload', ...)
```

**修改为**:
```javascript
// ✅ 改为 /admin/upload
formData.post('/admin/upload')
// 或
axios.post('/admin/upload', formData)
// 或
fetch('/admin/upload', ...)
```

#### 完整示例

<details>
<summary>点击展开完整代码示例</summary>

```vue
<template>
  <div class="upload-section">
    <div class="drop-zone" @drop="handleDrop" @dragover.prevent>
      <input 
        type="file" 
        ref="fileInput" 
        @change="handleFileSelect"
        multiple
      >
      <p>拖放文件到这里，或点击选择</p>
    </div>
    
    <div v-if="uploadProgress" class="progress">
      <div class="progress-bar" :style="{width: uploadProgress + '%'}"></div>
      <span>{{ uploadProgress }}%</span>
    </div>
    
    <ul v-if="uploadedFiles.length" class="uploaded-files">
      <li v-for="file in uploadedFiles" :key="file.id">
        {{ file.originalName }}
        <a :href="file.url" target="_blank">下载</a>
      </li>
    </ul>
  </div>
</template>

<script>
import { post } from '@/api/http'

export default {
  name: 'Upload',
  data() {
    return {
      uploadProgress: 0,
      uploadedFiles: []
    }
  },
  methods: {
    async uploadFile(file) {
      try {
        const formData = new FormData()
        formData.append('file', file)
        
        // ✅ 改为 /admin/upload
        const response = await axios.post('/admin/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
          onUploadProgress: (event) => {
            this.uploadProgress = Math.round(
              (event.loaded / event.total) * 100
            )
          }
        })
        
        if (response.data.code === 200) {
          this.uploadedFiles.push({
            id: response.data.data.id,
            originalName: response.data.data.originalName,
            url: response.data.data.url
          })
          this.uploadProgress = 0
        } else {
          alert(`上传失败: ${response.data.msg}`)
        }
      } catch (error) {
        console.error('上传文件失败:', error)
        
        if (error.response?.data?.msg === '没有权限，仅管理员可以执行此操作') {
          alert('您没有权限上传文件，仅管理员可以上传')
        } else if (error.response?.data?.msg?.includes('文件过大')) {
          alert('文件大小不能超过 500MB')
        } else {
          alert('上传文件出错')
        }
      }
    },
    
    handleFileSelect(event) {
      const files = event.target.files
      for (let file of files) {
        this.uploadFile(file)
      }
    },
    
    handleDrop(event) {
      event.preventDefault()
      const files = event.dataTransfer.files
      for (let file of files) {
        this.uploadFile(file)
      }
    }
  }
}
</script>

<style scoped>
.drop-zone {
  border: 2px dashed #ccc;
  border-radius: 4px;
  padding: 40px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
}

.drop-zone:hover {
  border-color: #4CAF50;
  background-color: #f0f7f0;
}

.progress {
  margin-top: 20px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.progress-bar {
  height: 20px;
  background-color: #4CAF50;
  border-radius: 4px;
  flex: 1;
  transition: width 0.3s;
}

.uploaded-files {
  margin-top: 20px;
  list-style: none;
  padding: 0;
}

.uploaded-files li {
  padding: 10px;
  background-color: #f5f5f5;
  margin-bottom: 8px;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.uploaded-files a {
  color: #4CAF50;
  text-decoration: none;
}

.uploaded-files a:hover {
  text-decoration: underline;
}
</style>
```

</details>

---

## 🟢 可选改进 (推荐)

### 4. 添加权限检查

为了避免用户点击无权限按钮，可以在前端隐藏不该显示的功能。

#### 示例

```javascript
// 在 vue 组件中
computed: {
  isAdmin() {
    // 从 localStorage 或 Vuex 中读取用户角色
    const userRole = localStorage.getItem('userRole')
    return userRole === 'admin'
  }
}

// 在模板中
<button v-if="isAdmin" @click="createChallenge">创建题目</button>
<div v-else class="permission-denied">
  <p>您没有权限创建题目</p>
</div>
```

### 5. 改进错误处理

为了给用户更好的地验证，可以添加特定的错误处理：

```javascript
async createChallenge() {
  try {
    const response = await post('/admin/challenge', this.form)
    
    if (response.code === 403) {
      alert('您没有权限创建题目（仅管理员可以）')
      // 重定向到首页
      this.$router.push('/')
    } else if (response.code === 400) {
      alert(`参数错误: ${response.msg}`)
    } else if (response.code === 200) {
      alert('题目创建成功!')
      this.$router.push('/challenges')
    }
  } catch (error) {
    console.error(error)
    alert('请求失败，请检查网络连接')
  }
}
```

---

## 📋 修改检查清单

在部署前，请逐一检查：

- [ ] 📄 **CreateChallenge.vue**
  - [ ] 检查 POST endpoint 已改为 `/admin/challenge`
  - [ ] 验证代码中没有 `/challenge`（旧路径）
  - [ ] 确认添加了权限检查（可选）

- [ ] 📄 **CreateCompetition.vue**
  - [ ] 检查 POST endpoint 已改为 `/admin/competition`
  - [ ] 验证代码中没有 `/competitions`（旧路径）
  - [ ] ⭐ 确认添加了 `mode` 字段
  - [ ] 确认添加了时间验证

- [ ] 📄 **Upload API/Component**
  - [ ] 检查 POST endpoint 已改为 `/admin/upload`
  - [ ] 验证代码中没有 `/upload`（旧路径）
  - [ ] 确认添加了错误处理（>500MB 文件）

- [ ] 🧪 **本地测试**
  - [ ] 使用管理员账户登录
  - [ ] 测试创建题目（POST /admin/challenge）
  - [ ] 测试创建比赛（POST /admin/competition）
  - [ ] 测试上传文件（POST /admin/upload）
  - [ ] 使用普通账户尝试上述操作，应该收到 403 错误

---

## 🔍 快速查找和替换

如果你的代码编辑器支持查找和替换，可以使用以下快捷方法：

### VS Code 查找和替换

1. 按 `Ctrl+H` 打开查找和替换
2. 在左框输入要查找的字符串
3. 在右框输入替换内容
4. 点击"替换全部"

#### 替换列表

```
查找: post('/challenge',
替换: post('/admin/challenge',

查找: post('/challenges',
替换: post('/admin/challenge',

查找: post('/competitions',
替换: post('/admin/competition',

查找: post('/upload',
替换: post('/admin/upload',

查找: "/challenge"
替换: "/admin/challenge"

查找: "/competitions"
替换: "/admin/competition"

查找: "/upload"
替换: "/admin/upload"
```

---

## 📞 故障排查

### 问题 1: 403 Permission Denied

**错误信息**: `没有权限，仅管理员可以执行此操作`

**原因**: 当前用户不是管理员

**解决方案**:
```javascript
// 检查用户角色
console.log(localStorage.getItem('userRole'))

// 期望输出: 'admin'
// 实际输出: 'user' → 需要升级权限
```

### 问题 2: 404 Not Found

**原因**: Endpoint 路径错误

**检查清单**:
- [ ] 是否改为 `/admin/challenge`？
- [ ] 是否改为 `/admin/competition`？
- [ ] 是否改为 `/admin/upload`？
- [ ] 是否小写了 `/challenge` vs `/Challenge`？

### 问题 3: 400 Bad Request

**原因**: 缺少必要参数

**CreateCompetition 特别注意**:
```javascript
// ❌ 错误：缺少 mode 参数
{
  title: "...",
  start_time: "...",
  end_time: "..."
}

// ✅ 正确：包含 mode 参数
{
  title: "...",
  mode: 0,  // 必须！
  start_time: "...",
  end_time: "..."
}
```

### 问题 4: 413 Payload Too Large

**原因**: 上传文件超过 500MB

**解决方案**:
```javascript
// 在上传前检查文件大小
const MAX_SIZE = 500 * 1024 * 1024  // 500MB
if (file.size > MAX_SIZE) {
  alert('文件大小不能超过 500MB')
  return
}
```

---

## 📚 相关文档

- **后端实现指南**: `docs/IMPLEMENTATION_GUIDE.md`
- **权限系统设计**: `docs/admin_permission_system.md`
- **变更记录**: `docs/ADMIN_PERMISSION_CHANGELOG.md`

---

## ✅ 完成确认

修改完成后，请确认：

```
[ ] 所有 3 个文件已修改
[ ] 已移除所有旧的 endpoint 路径
[ ] 本地运行通过测试
[ ] 管理员可以创建题目、比赛、上传文件
[ ] 普通用户被拒绝访问管理员接口
```

完成以上步骤后，前端已经准备好与新的后端权限系统配合工作！🎉

