swagger: '2.0'

info:
  description: the server API of the Homeworks
  version: 1.0.0
  title: server API
    
host: 47.102.120.167:2333
basePath: /api
tags:
  - name: user
    description: user
  - name: day
    description: day
  - name: night
    description: night

 
    
schemes:
  - https
  - http
  
paths:
  /user/login/:
    post:
      tags:
        - user
      summary: 登录
      description: login for account
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: data
          in: body
          description: post data
          schema:
            required:
              - Sid
              - Password
            properties:
              Sid:
                type: string
                description: 学号
              Password:
                type: string
                description: 密码
      responses:
        200:
          description: 登录成功
          schema:
            type: object
            required:
              - token
            properties:
              token:
                type: string
              msg:
                type: string
                description: 登录结果信息
        401:
          description: 用户名或密码错误
          schema:
            type: object
            properties:
              msg:
                type: string
   #   security:
    #    - petstore_auth:
     #       - 'write:pets'
      #      - 'read:pets'
      
  /day/requirement/history:
    get:
      tags:
        - day
      summary: 历史发布需求
      description: Get requirement history
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        200:
          description: successful operation
          schema:
            type: object
            properties:
              msg:
                type: string
              history:
                type: array
                items:
                  properties:
                    title:
                      type: string
                    content:
                      type: string
                    requirementTime:
                      type: string
                    timeForm:
                      type: integer
                    timeEnd:
                      type: integer
                    type:
                      type: integer
                    requiremPeopleNum:
                      type: integer
                    place:
                      type: integer
                    tag: 
                      type: integer
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
            
   #   security:
    #    - petstore_auth:
     #       - 'write:pets'
      #      - 'read:pets'
      
  /night/secret/history:
    get:
      tags:
        - night
      summary: 历史发布秘密 
      description: get history of secrets
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string 
      responses:
        200:
          description: successful operation
          schema:
            type: object
            properties:
              msg:
                type: string
              history:
                type: array
                items:
                  properties:
                    title:
                      type: string
                    content:
                      type: string
                    sendTime:
                      type: integer
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
                
                
      #该路由是否不推荐使用/不可用
      
  /day/remind/content/:
    get:
      tags:
        - day
      summary: 获取白天消息提醒 
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
              data:
                type: array
                items:
                  properties:
                    title:
                      type: string
                    nickname:
                      type: string
                    gender:
                      type: integer
                    college:
                      type: string
        400:
          description: 参数不全
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
      #security:
       # - petstore_auth:
        #    - 'write:pets'
         #   - 'read:pets'
  
  /day/message/confine:
    get:
      tags:
        - day
      summary: 确认是否接受申请 
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        '200':
          description: successful operation
          schema:
            type: object
            properties:
              msg:
                type: string
              title:
                type: string
              content:
                type: string
              requirementTime:
                type: string
              timeForm:
                type: integer
              timeEnd:
                type: integer
              type:
                type: integer
              requiremPeopleNum:
                type: integer
              place:
                type: integer
              tag: 
                type: integer
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string

  /user/homepage/:
    get:
      tags:
        - user
      summary: 获取个人信息
      description: user's information
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        200:
          description: successful operation
          schema:
            type: object
            properties:
              msg:
                type: string
              Sid:
                type: string
              Nickname:
                type: string
              College:
                type: string
              Gender:
                type: integer
              PersonalizedSignature:
                type: string
              ContactWay:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
                
  /night/remind/exist:
    get:
      tags:
        - night
      summary: 黑夜是否存在未读消息提醒    
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        200:
          description: success
          schema:
            properties:
              msg:
                type: string
              existence:
                type: boolean
        400:
          description: 参数不全
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
                      
  /night/remind/content:
    get:
      tags:
        - night
      summary: 获取黑夜消息提醒 
      parameters:
        - name: token
          in: header
          required: true
          type: string
      responses:
        200:
          description: successful operation
          schema:
            type: object
            properties:
              msg:
                type: string
              title:
                type: string
              content:
                type: string
              sendTime:
                type: integer
              commentdata:
                type: array
                items:
                  properties:
                    comment:
                      type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
            
  /user/changeInfo/:
    post :
      tags:
        - user
      summary: 修改个人信息
      description: change users' information
      consumes: 
        - application/json
      produces:
        - application/json
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: data
          in: body
          schema:
            required:
              - VerifyItem
              - VerifyInfo
            properties:
              VerifyItem:
                type: string
              VerifyInfo:
                type: string
      responses:
        200:
          description: change successfully
          schema:
            type: object
            properties:
              msg:
                type: string
        400:
          description: 修改失败
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
                
  /night/secret/create:
    post:
      tags:
        - night
      summary: 发布秘密
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: data
          in: body
          schema:
            required:
              - senderSid
              - title
              - content
              - sendTime
            properties:
              senderSid:
                type: string
              title:
                type: string
              content:
                type: string
              sendTime:
                type: integer
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
              secretId:
                type: integer
        400:
          description: 发布失败 
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string


  /night/comment/create:
    post:
      tags:
        - night
      summary: 发布评论 
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: data
          in: body
          schema:
            required:
              - senderSid
              - commentTime
              - content
              - secretId
            properties:
              senderSid:
                type: string
              commentTime:
                type: integer
              content:
                type: string
              secretId:
                type: integer
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
        400:
          description: 请求错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
 
  /day/requirement/create:
    post:
      tags:
        - day
      summary: 添加需求 
      parameters:
        - name: query 
          in: query
          required: true
          type: string
        - name: token
          in: header
          required: true
          type: string
        - name: data
          in: body
          schema:
            required:
              - senderSid
              - title
              - content
              - requirementTime
              - timeForm
              - timeEnd
              - type
              - requirePeopleNum
              - place
              - tag
            properties:
              senderSid:
                type: string
              title:
                type: string
              content:
                type: string
              requirementTime:
                type: string
              timeForm:
                type: integer
              timeEnd:
                type: integer
              type:
                type: integer
              requiremPeopleNum:
                type: integer
              place:
                type: integer
              tag: 
                type: integer
      responses:
        200:
          description: 发布需求成功 
          schema:
            type: object
            properties:
              msg:
                type: string
        400:
          description: 发布错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string


  /day/remind/exist:
    get:
      tags:
        - day
      summary: 白天是否存在未读消息提醒
      parameters:
        - name: token
          in: header
          required: true
          type: string        
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
              exsitence:
                type: boolean
        400:
          description: 请求错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
       
                
  /day/requirementsSquare:
    get:
      tags:
        - day
      summary: 请求筛选
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: query
          in: query
          required:
            - type1
              description:学习/运动/娱乐  
            - space
            - dataTime
            - timeForm
            - timeEnd
            - type2
              description:活动具体形式
          properties:
              type1:
                type: integer
                description: 初始值type1=1 
              space:
                type: string
                description: 若不选为零，其他选项依次从零递增，选择的组成字符串     
              dateTime:
                type: string
                description: 周一到周日选择了的为1，未选择的为0，组成字符串  
              timeForm:
                type: string
              timeEnd:
                type: integer
              type2:
                type: integer
              type:
                type: string
      responses:
        200:
          description: 成功 
          schema:
            type: object
            properties:
              msg:
                type: string
              data:
                type: array
                items:
                  properties:
                    title:
                      type: string
                    content:
                      type: string
                    equirementTime:
                      type: string
                    timeForm:
                      type: integer
                    timeEnd:
                      type: integer
                    type:
                      type: integer
                    requiremPeopleNum:
                      type: integer
                    place:
                      type: integer
                    tag: 
                      type: integer
                
        400:
          description: 请求错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string

  /mail/noticeTime/{noticeTimeId}/changeStatus/:
    put:
      tags:
        - mail
      summary: 改变时间节点启用状态
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: noticeTimeId
          in: path
          required: true
          type: string          
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
              statusMessage:
                type: string

        400:
          description: 请求错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
        404:
          description: 未找到
          schema:
            type: object
            properties:
              msg:
                type: string

  /mail/noticeTime/delete/:
    delete:
      tags:
        - mail
      summary: 移除时间节点
      parameters:
        - name: token
          in: header
          required: true
          type: string
        - name: data
          in: body
          schema:
            required:
              - noticeTimeId
            properties:
              noticeTimeId:
                type: string
      responses:
        200:
          description: success
          schema:
            type: object
            properties:
              msg:
                type: string
        400:
          description: 请求错误
          schema:
            type: object
            properties:
              msg:
                type: string
        401:
          description: 身份认证错误
          schema:
            type: object
            properties:
              msg:
                type: string
        404:
          description: 未找到
          schema:
            type: object
            properties:
              msg:
                type: string  
