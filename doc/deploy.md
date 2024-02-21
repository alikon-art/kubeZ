# 部署文档

### 部署后端
创建文件kubez.yaml,然后kubectl apply -f 它
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubez-configmap
  namespace: default
data:
  config.yaml: |-
    logs:
      level: debug
      timeformat: 2006/1/02 15:04:05
      showcaller: false
    program:
      # 后端程序运行的端口号
      port: 8080
      # 
      jwtsecret: jwtHello
      username: admin
      password: 123
      kubeconfig: ===== ## 你的k8s集群的kubeconfig,base64格式
      
immutable: false
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubez
  labels:
    app: kubez
  annotations:
    app: kubez
  namespace: default
spec:
  selector:
    matchLabels:
      app: kubez
  replicas: 1
  template:
    metadata:
      labels:
        app: kubez
      annotations:
        app: kubez
    spec:
      affinity:
        podAntiAffinity: {}
      restartPolicy: Always
      imagePullSecrets: []
      dnsPolicy: ClusterFirst
      hostNetwork: false
      volumes:
        - name: kubez-config
          configMap:
            name: kubez-configmap
            defaultMode: 420
            items: []
            optional: true
      containers:
        - name: kubez
          image: reactive0/kubez:0.0.1
          tty: false
          workingDir: ''
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              memory: 1024Mi
              cpu: 1
            requests:
              memory: 128Mi
              cpu: 100m
          ports: []
          lifecycle: {}
          volumeMounts:
            - name: kubez-config
              mountPath: /kubez/config.yaml
              subPath: config.yaml
              readonly: false
          env: []
          envFrom: []
      initContainers: []
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
--- 

apiVersion: v1
kind: Service
metadata:
  name: kubez
  namespace: default
  labels: {}
  annotations: {}
spec:
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
      name: kubez
  selector:
    app: kubez
  type: ClusterIP
  sessionAffinity: None
```


### 部署前端
同上
```
apiVersion: v1
kind: Service
metadata:
  name: kubez-frontend
  namespace: default
  labels:
    app: kubez-frontend
spec:
  ports:
    - name: kubez-frontend
      protocol: TCP
      port: 8080
      targetPort: 80
      nodePort: 30261
  selector:
    app: kubez-frontend
  type: NodePort
  sessionAffinity: None

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubez-frontend
  namespace: default
  labels:
    app: kubez-frontend
  annotations:
    app: kubez-frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubez-frontend
  template:
      labels:
        app: kubez-frontend
      annotations:
        app: kubez-frontend
    spec:
      containers:
        - name: kubez-frontend
          image: reactive0/kubez_frontend:0.0.1
          ports:
            - name: kubez-frontend
              containerPort: 80
              protocol: TCP
          imagePullPolicy: Always
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
```

之后访问nodeport对应的端口,再在集群管理中添加集群即可
